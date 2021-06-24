package main

import (
	"fmt"
	"time"
)

var dataMap = make(map[int]int)
var dataCh = make(chan int)

func main() {
	max := 10000
	t1 := time.Now().UnixNano()
	for i := 0; i < max; i++ {
		go modifyByChan(i)
	}
	chanServe(max)
	t2 := time.Now().UnixNano()
	fmt.Printf("len dataMap = %d, time = %vms", len(dataMap), (t2-t1)/1000000)
	// len dataMap = 10000, time = 70ms
}

func modifyByChan(i int) {
	dataCh <- i
}

func chanServe(max int) {
	for {
		i := <-dataCh
		dataMap[i] = i
		if len(dataMap) == max {
			return
		}
	}
}
