## go 的 Context
* 一直对 go 的 Context 一知半解，不了解其用途，因此在这里着重了解一下 go 语言的 Context
* 飞雪无情的一个博文对 go 的 Context 讲的比较易懂一些，所以就先从[这篇博文](https://www.flysnow.org/2017/05/12/go-in-action-go-context.html)开始吧

### 常用的并发控制
* [飞雪无情](ttps://www.flysnow.org/2017/05/12/go-in-action-go-context.html)博客中提到，常用的并发控制是通过 sync 包中的 WaitGroup 来实现的

```go
// 使用 sync 包中的 WaitGroup 实现协程的并发控制
// 其使用场景是：多个协程分别完成一整件事的一部分的工作，等待前部协程都完成之后，才算是完成了一整件事
func contextPart1() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(1 * time.Second)
		log.Println("协程 1号")
		wg.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("channel 2 is complete")
		wg.Done()
	}()
	wg.Wait()
	log.Println("所有协程都执行完成~")
}
```

### channel + select
* 以上这种场景主要是针对，可以自行结束的协程的并发控制。如果遇到的场景是协程并不会自动结束，该如何处理呢？
* 有一种方式，就是在任务协程中监听一个 channel，这个 channel 中一旦有数据（控制信号）就意味着对任务协程状态的控制

```go
// 通过通知协程结束的方式控制协程
func contextPart2() {
	stopCh := make(chan bool)
	go func() {
		for {
			select {
			case <-stopCh:
				fmt.Println("协程即将结束了")
				return
			default:
				fmt.Println("默认情况，继续执行...")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	fmt.Println("通知协程要结束了")
	stopCh <- true
	time.Sleep(5 * time.Second)
}
```

* 上面代码中，一开始创建了一个 stopCh 的通道，表示任务协程终止信号。任务协程中，通过 select 语句监听 stopCh 是否有数据，如果有数据，则实现协程任务结束操作，也就是 return。
* 以下是飞雪博客中对这总 channel + select 的方式的评价：
>这种chan+select的方式，是比较优雅的结束一个goroutine的方式，不过这种方式也有局限性，如果有很多goroutine都需要控制结束怎么办呢？如果这些goroutine又衍生了其他更多的goroutine怎么办呢？如果一层层的无穷尽的goroutine呢？这就非常复杂了，即使我们定义很多chan也很难解决这个问题，因为goroutine的关系链就导致了这种场景非常复杂。

### 使用 Context
* 因为会有 goroutine 中开启 goroutine 的情况，为了能够更优雅的管理 goroutine，go 中引入了 Context。将上面的例子，使用 Context 的方式重写：

```go
// 使用 Context 方式管理 goroutine
func contextPart3() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("任务完成，结束...")
				return
			default:
				log.Println("goroutine 继续执行...")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	log.Println("通知任务协程，要结束了")
	cancel()
	time.Sleep(5 * time.Second)
}
```

* 看起来，跟 channel+select 的方式差不多嘛。实际的情况是使用 Context 方式可以更方便地控制、跟踪 goroutine。
>`context.Background()` 返回一个空的 Context，这个空的 Context 一般用于整个 Context 树的根节点。然后我们使用 `context.WithCancel(parent)` 函数，创建一个可取消的子 Context，然后当作参数传给 goroutine 使用，这样就可以使用这个子 Context 跟踪这个 goroutine。

* 在 goroutine 中，使用 `<-ctx.Done()` 来判断是否要结束。如果有值，则对应的分支直接 return。如果没有值，则继续处理对应的任务。
* 在 channel+select 的方式中，我们可以向对应的 channel 发送数据表示停止信号，那么 Context 的方式如何发送信号呢？
* 就是位于上方代码中的 `cancel()` 调用。它是 `context.WithCancel` 返回的 `CancelFunc` 类型

### Context 控制多个 goroutine
* 因为实际的场景中是非常的复杂而多样化的，一定存在着多个 goroutine，针对多个 goroutine，Context 是如何处理的呢？

```go
// 使用 Context 控制多个 goroutine
func contextPart4() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "task 1")
	go watch(ctx, "task 2")
	go watch(ctx, "task 3")

	time.Sleep(10 * time.Second)
	log.Println("可以通知任务结束")
	cancel()
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			log.Println(name + " 任务即将要退出了...")
			return
		default:
			log.Println(name + " goroutine 继续处理任务中...")
			time.Sleep(2 * time.Second)
		}
	}
}
```

* 通过实际的运行，可以看到只需要调用一次 `cancel()` 就能控制多个 goroutine 

### Context 接口
* 通过查看[官方代码](https://golang.org/src/context/context.go)，可以看到 Context 的接口：

```go
// 一个 Context 可以跨 API 携带截止时间、取消信号以及其他值
//
// Context 的方法可以由多个 goroutine 同时调用
type Context interface {
	// Deadline 返回代表上下文任务应该被取消的截止时间。当没有设置截止时间时，Deadline 将返回 ok==false。对 Deadline 的连续调用将返回相同的结果。
	Deadline() (deadline time.Time, ok bool)

	// Done 返回一个 channel，该 channel 在该上下文的任务被取消时应该被关闭。如果无法取消这个上下文，Done 可能返回 nil。对 Done 的连续调用将返回相同的值
	Done() <-chan struct{}

	// 如果 Done 没有关闭，Err 将返回 nil。
	// 如果 Done 已关闭，Err 返回一个非 nil 错误，原因是：如果上下文已被取消，则返回 cancel；如果上下文的截止时间已过，则返回 deadline。
	// 在 Err 返回非 nil 错误之后。对 Err 的连续调用将返回相同的错误。
	Err() error

	// Value 返回针对 key 上下文的关联的值，如果没有与 key 关联的值，则返回 nil。使用相同的 key 连续调用 Value 将返回一样的结果。
	Value(key interface{}) interface{}
}
```

* 其中有 4 个方法，相关的注释已经翻译为中文，可以了解一下各个方法的作用。
* 在 context 包中，已经实现了 2 种 Context，分别是 Background 和 TODO。从包的官方文档中可以看到：

```go
var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

// Background 返回一个非 nil 的空 Context。它不会被取消，没有值，也没有截止时间。它通常用于 main 函数、初始化和测试
// 并作为传入请求的顶层的 Context
func Background() Context {
	return background
}

// TODO 返回一个非 nil 的空 Context。当不清楚要使用哪个 Context 或者还不能用 Context 时，代码应该使用 context.TODO 
// （因为周边函数还没有扩展到接收一个 Context 参数）
func TODO() Context {
	return todo
}
```

* 从定义中可以看到没有什么特别复杂的东西，但是我们需要注意一下 emptyCtx。包中对它的定义如下：

```go
type emptyCtx int
```

>emptyCtx 不会被取消，它没有值，也没有截止时间。它不是 `struct{}`，因为这种类型的变量必须有不同的地址

* 这就是为什么 background 和 todo 明明类型一样，却还 new 两次。下面看一下 emptyCtx 类型实现了哪些方法：

```go
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}

func (*emptyCtx) Err() error {
	return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
	return nil
}

func (e *emptyCtx) String() string {
	switch e {
	case background:
		return "context.Background"
	case todo:
		return "context.TODO"
	}
	return "unknown empty Context"
}
```

* 貌似大部分方法都是直接返回一个 nil
* 在上面的 contextPart3、contextPart4 中，我们都是在 main 中创建 goroutine，我们讲过，实际的场景中，会有在子 goroutine 中创建 goroutine 的情况，这种时候，改如何创建呢？难道是还是使用 `ctx, cancel := context.WithCancel(context.Background())` 创建吗？
* 在回答这个问题前，我们先了解一下 `With*` 系列的方法：

```go
// WithCancel 返回父级的 Done 通道的副本。无论一开始发生什么，当调用其返回的 cancel 函数或当关闭父级 Context 的 Done 通道时，都将关闭返回的上下文的 Done 通道。 
// 取消此 Context 将释放与其关联的资源，因此，代码应该在此上下文中运行的操作完成后立即调用 cancel。
func WithCancel(parent Context) (ctx Context, cancel CancelFunc); 

// WithCancel 将返回一个父级 context 的副本，该副本带有截止时间调整为不迟于 d。如果父级上下文的截止时间比 d 要早，
// 则 WithDeadline(parent, d) 在语义上等同于父级 context。当截止时间过期时，或当调用返回的 cancel 时，
// 或当父级 context 的 Done 通道被关闭时，返回的 context 的 Done 通道将关闭，以最先发生的情况为准。

// 取消此上下文将释放与其资源关联的资源，因此，在此 context 中运行的操作完成后应该立即调用 cancel。
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc);

// WithTimeout 将返回 `WithDeadline(parent, time.Now().Add(timeout))` 参数
// 取消这个 context将释放与它相关联的资源，所以代码中应该在这个上下文中运行的操作完成后立即调用cancel:
//
// 	func slowOperationWithTimeout(ctx context.Context) (Result, error) {
// 		ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
// 		defer cancel()  // releases resources if slowOperation completes before timeout elapses
// 		return slowOperation(ctx)
// 	}
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc);

// WithValue 返回父级的一个副本，其中与 key 关联的值是 val
// 只对传输进程和 API 的请求域数据使用 context Value，而不是将可选参数传递给函数的场景

// 提供的 key 必须是可比较的，并且不应该是 string 或者其他内置类型，这样可以避免使用 context 在包之间发生冲突。
// 使用 WithValue 的用户应该为 key 定义自己的类型。为了避免在分配时分配给 interface{}，context key 通常有具体的类型 struct{}
// 或者，导出的上下文关 key 变量的静态类型应该是一个指针或者 interface。
func WithValue(parent Context, key, val interface{}) Context;
```

* 这 4 个函数，参数中都有 `parent Context`，也就是父级 Context，要达到“在子 goroutine 中创建 goroutine”，其实就是基于这个“父级 Context”来创建。可以将其称之为“衍生”
* 从抽象的角度看，“子 goroutine 中创建 goroutine ”达到一定程度可以将总体看成一颗 Context 树
>树的每个节点都可以有任意多个子节点，节点层级可以有任意多个

* 我们可以观察到，`With*` 系列函数中的前三个都会返回 `CancelFunc` 类型：

```go
// CancelFunc 的主要作用就是放弃其任务的操作。CancelFunc 不会等待其任务停止
// 在第一次调用 CancelFunc 后，后续的调用将什么都不会做。也就是只会生效一次。
type CancelFunc func()
```

### WithValue
* 由于在“子 goroutine 中创建 goroutine ”的过程中，我们可能需要在比较深的 goroutine 中需要使用外层就产生的一些数据，此时我们可以使用 WithValue
* WithValue 可以帮我们传递一些必须的元数据，这些数据会附加在 Context 中

```go
func contextPart5() {
	var key string = "key1"
	ctx, cancel := context.WithCancel(context.Background())
	// 附加的数据
	vCtx := context.WithValue(ctx, key, "这里是元数据-任务1")
	go watch2(vCtx, "task1")
	time.Sleep(10 * time.Second)
	log.Println("可以通知任务停止了...")
	cancel()
	time.Sleep(5 * time.Second)
}

func watch2(ctx context.Context, name string) {
	var key string = "key1"
	for {
		select {
		case <-ctx.Done():
			log.Println("获取到元数据：" + ctx.Value(key).(string))
			log.Println(name + " 任务即将要退出了...")
			return
		default:
			log.Println(name + " goroutine 继续处理任务中...")
			time.Sleep(2 * time.Second)
		}
	}
}
```

* 通过运行并观察，输出如下：

```other
2019/07/01 17:22:54 task1 goroutine 继续处理任务中...
2019/07/01 17:22:56 task1 goroutine 继续处理任务中...
2019/07/01 17:22:58 task1 goroutine 继续处理任务中...
2019/07/01 17:23:00 task1 goroutine 继续处理任务中...
2019/07/01 17:23:02 task1 goroutine 继续处理任务中...
2019/07/01 17:23:04 可以通知任务停止了...
2019/07/01 17:23:04 获取到元数据：这里是元数据-任务1
2019/07/01 17:23:04 task1 任务即将要退出了...
```

* 内存的 goroutine 可以正确通过 context 获取到对应的元数据。值的注意的是，这里的值必须是线程安全的。

### Context 使用原则
* 以下是飞雪无情博客中提到的一些使用原则，为了能够更好、更准确的使用 Context，我们最好遵循：
>1.不要把Context放在结构体中，要以参数的方式传递
>2.以Context作为参数的函数方法，应该把Context作为第一个参数，放在第一位。
>3.给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO
>4.Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递
>5.Context是线程安全的，可以放心的在多个goroutine中传递

## 参考资料
* Go语言实战笔记（二十）| Go Context https://www.flysnow.org/2017/05/12/go-in-action-go-context.html
* 关于 Context 的另一篇文章 [https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39](https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39)
