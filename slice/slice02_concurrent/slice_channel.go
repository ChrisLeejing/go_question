package main

import (
	"fmt"
	"time"
)

var list3 = []int{}
var chanList = make(chan int)

func main() {
	max := 10000
	t1 := time.Now().UnixNano()
	for i := 0; i < max; i++ {
		go addNotSafe(i)
	}
	sliceServe()
	t2 := time.Now().UnixNano()
	fmt.Printf("slice channel safe: len list3 = %d, time = %dms\n", len(list3), (t2-t1)/1000000)
	// slice channel safe: len list3 = 10000, time = 64ms
}

func addNotSafe(i int) {
	chanList <- i
}

func sliceServe() {
	for {
		data := <-chanList
		list3 = append(list3, data)
		if len(list3) == 10000 {
			return
		}
	}
}
