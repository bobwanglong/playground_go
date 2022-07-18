### 观察者模式
demo待完善
# go观察者模式

### 一、概念

观察者模式(Observer Pattern)，定义对象间的一种一对多依赖关系，使得每当一个对象状态发生改变时，其相关依赖对象皆得到通知后，可自行调节用自身的处理程序，实现想要干的事情，比如更新自己的状态。

发布者对观察者唯一了解的是它实现了某个接口(观察者接口)。这种松散耦合的设计最大限度地减少了对象之间的相互依赖，因此使我们可以构建灵活的系统。

### 二、自己理解

观察者模式也经常被叫做发布-订阅(Pulish/Subscribe)模式，概念中定义对象间的一种一对多依赖关系，“一”指的是发布变更的主体对象，“多”，指的是订阅变更通知的订阅者对象

发布的状态变更信息会被包装到一个对象里，这个对象被称为事件，事件一般用英语过去式的语态来命名，比如用户注册时，用户模块在用户创建好后发布一个时间UserCreated或者UserWascreated都行，这样从名字上就能看得出来是一个已经发生过的事件了

时间发布给订阅者的过程，其实就是便利一下已经注册的事件订阅者，逐个去调用订阅者实现的观察者接口方法，比如叫 handleEvebt之类的方法，这个方法的参数一般就是当前的事件对象。

至于很多人会好奇的，时间的处理是不是异步的？主要看具体的需求，一般是同步的，即发布事件后，触发事件的方法会阻塞等到全部订阅者返回后再继续，当然也可以让订阅者的处理异步执行，完全取决于需求

大部分场景下其实是同步执行的，单体架构会在一个数据库事务里持久化因为主体状态变更，而需要更改的所有实体类。

微服务架构下常见的做法是有一个事件储存，订阅者接到事件通知后，会把事件先存在到事件驱存储中，这量布也需要在一个事务里完成才能保证最终一致性，后面会再有其他线程把事件从事件存储中搞到消息设施里，发给其他服务。从而在微服务架构下实现各个位于不同服务的实体间的最终一致性

所以，从效率上看，观察者模式在大多数下没啥提升，更多的是达到一种程序节后上的解耦，让代码不至于那么难维护

### Go实现观察者模式

demo

```go

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
```