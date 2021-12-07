package main

import (
	"fmt"
	// "playground/myatomic"
	"github.com/logrusorgru/aurora"
)

func main() {
	fmt.Println("hello playground")

	// fmt.Println("hello atomic")
	// myatomic.Run()
	a := aurora.NewAurora(true)
	fmt.Println(a.Red("Red"))
	fmt.Println(aurora.Green("aaa"))
}
