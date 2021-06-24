package main

import (
	"fmt"
	"sync"
	"time"
)

var list1 = []int{}
var list2 = []int{}
var wg = sync.WaitGroup{}
var mutex = sync.Mutex{}

func main() {
	max := 10000
	t1 := time.Now().UnixNano()
	wg.Add(max)
	for i := 0; i < max; i++ {
		go appendNotSafe(i)
	}
	wg.Wait()
	t2 := time.Now().UnixNano()
	fmt.Printf("slice not safe: len list1 = %d, time = %dms\n", len(list1), (t2-t1)/1000000)

	t3 := time.Now().UnixNano()
	wg.Add(max)
	for i := 0; i < max; i++ {
		go appendSafe(i)
	}
	wg.Wait()
	t4 := time.Now().UnixNano()
	fmt.Printf("slice safe: len list2 = %d, time = %dms\n", len(list2), (t4-t3)/1000000)
	// slice not safe: len list1 = 8354, time = 5ms
	// slice safe: len list2 = 10000, time = 2ms
}

func appendNotSafe(i int) {
	list1 = append(list1, i)
	wg.Done()
}

func appendSafe(i int) {
	mutex.Lock()
	list2 = append(list2, i)
	mutex.Unlock()
	wg.Done()
}
