package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "aaabbb"
	prefix := "aaa"

	res := strings.TrimPrefix(s, prefix)
	fmt.Println(res)
	// 判断是否含有中文
	chineseStr := "I love 中国"
	fmt.Println(isChinese(chineseStr))
}
func isChinese(s string) bool {
	// var count int
	for _, v := range s {
		if unicode.Is(unicode.Han, v) {
			return true
		}
	}
	return false
}
