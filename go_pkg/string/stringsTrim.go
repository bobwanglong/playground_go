package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aaabbb"
	prefix := "aaa"

	res := strings.TrimPrefix(s, prefix)
	fmt.Println(res)
}
