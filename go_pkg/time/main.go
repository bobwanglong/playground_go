package main

import (
	"fmt"
	"time"
)

func main() {
	old := time.Now()
	time.Sleep(time.Second)
	// duration := old.Sub(time.Now())
	duration := time.Until(old) // t.sub的缩写，表示过去一段时间持续了多久

	fmt.Println(duration)
}
