package main

import "fmt"

type A interface {
	B
	C
}
type B interface {
	bar()
}
type C interface {
	foo()
}

func do(a A) {
	a.bar()
}

type APP struct{}

func (app *APP) bar() {
	fmt.Println("app bar")
}
func (app *APP) foo() {
	fmt.Println("app foo")
}

func cDo(c C) {
	fmt.Println("c interface")
}
func main() {
	app := &APP{}
	// app.bar()
	cDo(app)
	do(app)
}
