package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Second * 2)
	go func() {
		for t := range t.C {
			fmt.Println(t.Format("2006-01-02 15:04:05"))
			p()
			fmt.Println("hello I'ok")
		}
	}()

	time.Sleep(time.Second * 10)
	t.Stop()
}

func p() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover")
		}
	}()
	panic("wrong")
}
