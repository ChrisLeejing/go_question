package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	// 场景1：正常输出
	// channel buffer data: 1
	// channel buffer data: 2
	// channel buffer data: 3
	// channel buffer data: 4
	go fun2(ch)
	ch <- 4

	// 场景2：阻塞抛错：fatal error: all goroutines are asleep - deadlock!
	// ch <- 4
	// go fun2(ch)

	time.Sleep(time.Second)
}

func fun2(ch chan int) {
	for {
		select {
		case i := <-ch:
			fmt.Printf("channel buffer data: %d\n", i)
		}
	}
}
