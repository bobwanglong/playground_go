package main

import "fmt"

func main() {
	sub := &SubjectImpl{}

	sub.Subscribe(&Observer1{})
	sub.Subscribe(&Observer2{})
	sub.Notify("hello")

}

type Subject interface {
	Subscribe(observer Observer)
	Notify(msg string)
}
type Observer interface {
	Update(msg string)
}

type SubjectImpl struct {
	observer []Observer
}

func (sub *SubjectImpl) Subscribe(observer Observer) {
	sub.observer = append(sub.observer, observer)
}

func (sub *SubjectImpl) Notify(msg string) {
	for _, o := range sub.observer {
		o.Update(msg)
	}
}

type Observer1 struct{}

func (Observer1) Update(msg string) {
	fmt.Printf("observer1: %s\n", msg)
}

type Observer2 struct{}

func (Observer2) Update(msg string) {
	fmt.Printf("observer2: %s\n", msg)
}
