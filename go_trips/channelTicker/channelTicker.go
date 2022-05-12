package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")
	finishCh := make(chan struct{})
	mainCh := make(chan struct{})
	go do(finishCh, mainCh)
	time.Sleep(time.Second * 10)
	finishCh <- struct{}{}
	<-mainCh
	fmt.Println("主程序退出")
}

func do(finishCh, mainCh chan struct{}) {
	ticker := time.NewTicker(time.Second * 2)
	defer ticker.Stop()
	num := 0
	for {
		select {
		case <-finishCh:
			fmt.Println("子goroutine finished")
			mainCh <- struct{}{}
			return
		case <-ticker.C:
			num += 1
			fmt.Println("num:", num)
		}
	}
}
