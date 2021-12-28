package main

import (
	"fmt"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println("show A")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("show B")
}
func (p People) ShowC() {
	fmt.Println("show C")
}
func (p *People) ShowD() {
	fmt.Println("show D")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher show B")
}

type A int

func (a A) Value() int {
	return int(a)
}

func (a *A) Set(n int) {
	*a = A(n)
}

type B struct {
	A
	b int
}

type C struct {
	*A
	c int
}

func main() {
	// var t Teacher
	t := Teacher{}
	t.ShowA() // show A show B
	t.ShowC() // show C
	t.ShowD()
}
