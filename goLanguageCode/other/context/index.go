package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctxa, cancel := context.WithCancel(context.Background())
	go work(ctxa, "work1")
	fmt.Println(cancel)
	time.Sleep(3 * time.Second)
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "get message to quit")
			return
		default:
			fmt.Println(name, "is running")
			time.Sleep(time.Second)
		}
	}
}
