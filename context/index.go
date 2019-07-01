package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	// contextPart1()
	// contextPart2()
	// contextPart3()
	// Context 控制多个 goroutine
	// contextPart4()
	// 使用 WithValue
	contextPart5()
}

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
