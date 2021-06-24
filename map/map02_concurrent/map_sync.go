package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Map

func main() {
	// 写入
	data := []string{"a", "b", "c", "d"}
	for i := 0; i < 4; i++ {
		go func(i int) {
			m.Store(i, data[i])
		}(i)
	}
	time.Sleep(time.Second)

	// 读取
	v, ok := m.Load(0)
	fmt.Printf("Load: %v, %v\n", v, ok)

	// 删除
	m.Delete(1)

	// 读或写
	v, ok = m.LoadOrStore(1, "bb")
	fmt.Printf("LoadOrStore: %v, %v\n", v, ok)

	// 遍历
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("Range: %v, %v\n", key, value)
		return true
	})
}
