package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var m sync.Map
	paths := []string{"/hello", "/hello", "/hello/"}
	for _, v := range paths {
		if value, ok := m.Load(v); ok {
			v1 := value.(int64)
			fmt.Println(v1)
			value = atomic.AddInt64(&v1, 1)
			fmt.Println(v1)
			m.Store(v, value)
		} else {
			m.Store(v, int64(1))
		}

	}
	m.Range(func(key, value interface{}) bool {
		path := key.(string)
		num := value.(int64)
		fmt.Println(path, num)
		return true
	})

}
