package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	kafka1 := "192.168.10.100:49093"
	kafka2 := "192.168.10.92:49092"
	kafka3 := "192.168.10.91:49094"
	client, err := sarama.NewClient([]string{kafka1, kafka2, kafka3}, config)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// get partitionId list
	partitions, err := consumer.Partitions("test")
	if err != nil {
		panic(err)
	}

	for _, partitionId := range partitions {
		// create partitionConsumer for every partitionId
		partitionConsumer, err := consumer.ConsumePartition("test", partitionId, sarama.OffsetOldest)
		if err != nil {
			panic(err)
		}

		go func(pc *sarama.PartitionConsumer) {
			defer (*pc).Close()
			// block
			for message := range (*pc).Messages() {
				value := string(message.Value)
				log.Printf("Partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, value)
			}

		}(&partitionConsumer)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	// select {
	// case <-signals:

	// }
	<-signals
}
