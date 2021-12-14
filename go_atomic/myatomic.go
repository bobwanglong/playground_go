package myatomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Rectangle struct {
	lengh int
	width int
}

var rect atomic.Value
var wg sync.WaitGroup
var rectLocal = new(Rectangle)

func update(width, lengh int) {
	// rectLocal := new(Rectangle)
	rectLocal.lengh = lengh
	rectLocal.width = width
	// rect.Store(rectLocal)

}

func Run() {
	wg.Add(10)
	count := 10
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			update(i, i+5)
		}()
	}
	wg.Wait()
	// _r := rect.Load().(*Rectangle)
	// fmt.Printf("rect.width=%d\nrect.lenth=%d\n", _r.width, _r.lengh)	// _r := rect.Load().(*Rectangle)
	fmt.Printf("rect.width=%d\nrect.lenth=%d\n", rectLocal.width, rectLocal.lengh)

}
