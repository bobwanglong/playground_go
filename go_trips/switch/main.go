package main

import "fmt"

func F() bool {
	return false
}

func main() {
	switch F(); {
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	}
}
