package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// 基于sarama第三方库开发的kafka client

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test"
	strMsg := "hello"
	// strMsg := `{"schema":"publish",  "appGroupId":"2123","appId":"2230", "apiName":"loginfo", "host":"192.168.24.212:8000", "authMethod":"","requestProtocal":"http", "apiPath":"/loginfo", "apiId": "912394844830171138"}`
	// strMsg := `{"schema":"update",  "appGroupId":"2123","appId":"2230", "apiName":"loginfo", "host":"192.168.24.212:8001", "authMethod":"","requestProtocal":"http", "apiPath":"/loginfo", "apiId": "912394844830171138"}`

	// strMsg := `{"schema":"remove",  "appGroupId":"2123","appId":"2230", "apiName":"loginfo"}`
	// strMsg := `{"schema":"update",  "appGroupId":"54","appId":"251", "apiName":"hello", "host":"192.168.24.212:8000", "authMethod":"","requestProtocal":"http", "apiPath":"/hello/{id}", "apiId": "908717929749544960"}`

	// strMsg1 := `{"schema":"remove",  “apiName":"findUsingPOST"}`
	msg.Value = sarama.StringEncoder(strMsg)
	// 连接kafka
	kafka1 := "192.168.10.100:49093"
	kafka2 := "192.168.10.92:49092"
	kafka3 := "192.168.10.91:49094"

	// kafka := "192.168.10.90:49092"

	client, err := sarama.NewSyncProducer([]string{kafka1, kafka2, kafka3}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
