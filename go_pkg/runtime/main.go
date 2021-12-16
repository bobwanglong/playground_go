// package main

// import "runtime"

// func f() {
// 	defer runtime.Goexit()
// 	panic("bye")
// }

// func main() {
// 	go f()

// 	// 等待直到f协程退出。
// 	for runtime.NumGoroutine() > 1 {
// 		runtime.Gosched()
// 	}
// }
package main

import (
	"fmt"
	"runtime"
)

func f() {
	defer func() {
		recover()
	}()
	defer panic(123)
	runtime.Goexit()
}

func main() {
	fmt.Println(runtime.NumCPU())     // 获取cpu的个数
	fmt.Println(runtime.NumCgoCall()) // 调用cgo的次数
	go func() {
		f()
		for {
			runtime.Gosched()
		}
	}()

	for runtime.NumGoroutine() > 1 {
		runtime.Gosched()
	}
}
