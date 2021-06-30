package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	// 场景1：fatal error: all goroutines are asleep - deadlock! 发生阻塞错误。
	// ch <- 1
	// go fun1(ch)

	// 场景2：channel unbuffer data: 1
	go fun1(ch)
	ch <- 1

	time.Sleep(time.Second)
}

func fun1(ch chan int) {
	for {
		select {
		case i := <-ch:
			fmt.Printf("channel unbuffer data: %d", i)
		}
	}
}
