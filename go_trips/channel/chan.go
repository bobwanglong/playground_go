// package main

// import (
// 	"fmt"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"
// )

// func main() {
// 	// 创建一个sig的channel，捕获系统的信号，传递到sig中
// 	sig := make(chan os.Signal)
// 	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)
// 	// 利用两个channel实现优雅退出
// 	// 父goroutine通知子goroutine准备优雅的关掉，也就是 stopCh
// 	// 子goroutine用纸父goroutine已经完成关闭，也是酒finishedCh
// 	stopCh := make(chan struct{})
// 	finishedCh := make(chan struct{})

// 	// mux := http.NewServeMux()
// 	// mux.HandleFunc("/hello", hello)

// 	// go http.ListenAndServe(":8888", mux) // 异步服务
// 	go func(stopCh, finishedCh chan struct{}) {
// 		for {
// 			select {
// 			case <-stopCh:
// 				fmt.Println("stpped")
// 				finishedCh <- struct{}{}
// 				return
// 			default:
// 				time.Sleep(time.Second)
// 			}
// 		}
// 	}(stopCh, finishedCh)

// 	fmt.Println(<-sig)
// 	stopCh <- struct{}{}
// 	<-finishedCh
// 	fmt.Println("finished")

// }

// // func hello(w http.ResponseWriter, r *http.Request) {
// // 	fmt.Println("hello")
// // }

// // 华丽的解决方案 - channel嵌套channel

// package main

// import (
// 	"fmt"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"
// )

// func main() {
// 	sig := make(chan os.Signal)
// 	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

// 	stopCh := make(chan chan struct{})

// 	go func(stopCh chan chan struct{}) {
// 		for {
// 			select {
// 			case ch := <-stopCh:
// 				fmt.Println("stopped")
// 				ch <- struct{}{}
// 				return
// 			default:
// 				time.Sleep(time.Second)
// 			}
// 		}
// 	}(stopCh)
// 	<-sig

// 	ch := make(chan struct{})
// 	stopCh <- ch
// 	<-ch
// 	fmt.Println("finished")
// }

// 标准解决方案
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	ctx, cancel := context.WithCancel(context.Background())
	finishCh := make(chan struct{})

	go func(ctx context.Context, finishCh chan struct{}) {
		for {
			select {
			case <-ctx.Done():
				// 结束后，通过ch通知主goroutine
				fmt.Println("stopped")
				finishCh <- struct{}{}
				return
			default:
				time.Sleep(time.Second)
			}
		}

	}(ctx, finishCh)

	<-sig
	// 收到结束信号，通知子goroutine结束程序，资源回收
	cancel()
	<-finishCh // 主程序结束
	fmt.Println("finished")
}
