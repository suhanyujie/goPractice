## 第4章 并发编程

### 优点
* 并发能更客观的表现问题模型
* 并发可以充分利用CPU核心的优势，提高程序的执行效率
* 并发能充分利用CPU与其他硬件设备固有的异步性
* Go 语言在语言级别支持轻量级线程，叫goroutine。Go 语言标准库提供的所有系统调用操作
  （当然也包括所有同步 IO 操作），都会出让 CPU 给其他goroutine。这让事情变得非常简单，让轻
  量级线程的切换管理不依赖于系统的线程和进程，也不依赖于CPU的核心数量。
* 不要通过共享内存来通信，而应该通过通信来共享内存

### select
* select的用法与switch语言非常类似，由select开始一个新的选择块，每个选择条件由
  case语句来描述。与switch语句可以选择任何可使用相等比较的条件相比，select有比较多的
  限制，其中最大的一条限制就是每个case语句里必须是一个IO操作

### 单向channel
* 它可以起到限定的作用
* 例如，c语言中函数声明时，针对指针形参加上const限定，这样函数体中就不允许修改对应的变量
* 单向通道的作用也和这个类似
* 它的定义如下：
>ch5 := <-chan int(ch4) // ch5就是一个单向的读取channel<br>
 ch6 := chan<- int(ch4) // ch6 是一个单向的写入channel
 
### 关闭channel
* 直接使用close函数
* 如果重复关闭一个通道，会产生panic
* 因此，读取的时候，可以使用多重返回值

```
x,ok := <-ch
```

* 如果ok返回false，则表示ch已经关闭

### 设置CPU核心数实现并行计算
* `runtime.GOMAXPROCS(3)`

### 同步锁
* go语言包中提供了sync包，包含2中锁类型：
#### sync.Mutex
* 当一个goroutine获得了Mutex后，其他goroutine就只能乖乖等到这个goroutine释放该Mutex

####sync.RWMutex
* 单写多读模型。阻止写，但不阻止读

### 原子性 全局唯一性操作
* 可以利用go提供的`sync.Once`来防止并发问题
* 除此之外，可以利用sync提供的`atomic`包









