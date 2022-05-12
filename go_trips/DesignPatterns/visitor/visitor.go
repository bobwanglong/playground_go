package main

// 定义访问的函数类型
type VisitorFunc func(*Info, error) error

// Visitor接口设计
type Visitor interface {
	Visit(VisitorFunc) error
}

// 资源对象
type Info struct {
	NameSpace   string
	Name        string
	OtherThings string
}

// 将Visitor函数应用到资源对象上
func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info, nil)
}

type NameVisitor struct {
	visitor Visitor
}

func (v NameVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(i *Info, e error) error {
		e = fn(i, e)
		//
		return e
	})
}
func main() {}
