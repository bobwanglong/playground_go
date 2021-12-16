package main

import "fmt"

func main() {
	f1()
	f2()
	f3()
}

type T struct {
	n int
}

func f1() {
	ts := [2]T{}
	for i, t := range ts {
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ") // 0
		}
	}
	fmt.Println(ts) // [{0} {9}]

}
func f2() {
	ts := [2]T{}
	for i, t := range &ts {
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
		}
	}
	fmt.Println(ts) // 9 [{0} {9}]

}

func f3() {
	ts := [2]T{}
	for i := range ts[:] {
		t := &ts[i]
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
		}
	}
	fmt.Println(ts) //9 [{3} {9}]
}
