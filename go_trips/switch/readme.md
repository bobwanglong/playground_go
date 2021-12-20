### switch流程控制代码块里的switch表达式总是被估值为类型确定值。例如，在下列switch代码块中的switch表达式123被视为一个int值，而不是一个类型不确定的整数。

```go
func main() {
	switch 123 {
	case int64(123): // error: 类型不匹配
	case uint32(789): // error: 类型不匹配
	}
}
```
### switch bool 
```go

func F() bool {
	return false
}

func main() {
	switch F(); {
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	}
}
```
等价于
```go
func main() {
	switch F(); true{
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	}
}
```