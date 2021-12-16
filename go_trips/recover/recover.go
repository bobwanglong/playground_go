package main

import "fmt"

func main() {
	p()
	fmt.Println("hello I'ok")
}

func p() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover")
		}
	}()
	panic("wrong")
}
