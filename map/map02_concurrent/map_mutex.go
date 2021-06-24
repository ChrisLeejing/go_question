package main

import (
	"fmt"
	"sync"
	"time"
)

var data = make(map[int]int)
var wgMap = sync.WaitGroup{}
var muMap = sync.Mutex{}

func main() {
	max := 10000

	wgMap.Add(max)
	t1 := time.Now().UnixNano()

	for i := 0; i < max; i++ {
		go modifySafe(i)
	}
	wgMap.Wait()
	t2 := time.Now().UnixNano()

	fmt.Printf("data len = %d, time=%dms", len(data), (t2-t1)/1000000)
	// data len = 10000, time=7ms
}

func modifySafe(i int) {
	muMap.Lock()
	data[i] = i
	muMap.Unlock()
	wgMap.Done()
}
