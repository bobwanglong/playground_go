package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	c := make(chan int)
	switch { // 缺省的bool 是true
	case true:
		run2(c)
	case false:
		run1(c)
	}

	fmt.Println("run...")
	<-c
}

func osExit() {
	os.Exit(24)
}
func run1(c chan int) {
	go func() {
		defer func() { c <- 1 }()
		defer fmt.Println("Go")
		func() {
			defer fmt.Println("C")
			runtime.Goexit() // 退出 所在的goroutine
		}()
		// 不会被执行到
		fmt.Println("Java")
	}()
}
func run2(c chan int) {
	go func() {
		defer func() { c <- 1 }()
		defer fmt.Println("Go")
		osExit() // 程序退出，什么都不会打印
		func() {
			defer fmt.Println("C")
			runtime.Goexit() // 退出 所在的goroutine
		}()
		// 不会被执行到
		fmt.Println("Java")
	}()
}
