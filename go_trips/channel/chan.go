package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 创建一个sig的channel，捕获系统的信号，传递到sig中
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)
	// 利用两个channel实现优雅退出
	// 父goroutine通知子goroutine准备优雅的关掉，也就是 stopCh
	// 子goroutine用纸父goroutine已经完成关闭，也是酒finishedCh
	stopCh := make(chan struct{})
	finishedCh := make(chan struct{})

	// mux := http.NewServeMux()
	// mux.HandleFunc("/hello", hello)

	// go http.ListenAndServe(":8888", mux) // 异步服务
	go func(stopCh, finishedCh chan struct{}) {
		for {
			select {
			case <-stopCh:
				fmt.Println("stpped")
				finishedCh <- struct{}{}
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}(stopCh, finishedCh)

	fmt.Println(<-sig)
	stopCh <- struct{}{}
	<-finishedCh
	fmt.Println("finished")

}

// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("hello")
// }
