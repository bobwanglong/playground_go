package main

import "fmt"

func decoration(f func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("decoration")
		f(s)
		fmt.Println("done")
	}
}
func foo(s string) {
	fmt.Printf("%s is string\n", s)
}
func main() {
	bar := decoration(foo)
	bar("abc")

}
