# 装饰模式

### 核心要旨

利用 *闭包* 实现, 一句话解释：在函数f前后，添加装饰性的功能函数，但不改变函数本身的行为。

### 场景

这种设计模式，对一些被高频使用的代码非常有用：

1. HTTP Server 被调用的hander
2. HTTP Client发送请求
3. 对MYSQL的操作

装饰性的功能，常见的有：

1. 打印相关的日志信息(Debug中非常有用)
2. 耗时相关的计算？
3. 监控埋点