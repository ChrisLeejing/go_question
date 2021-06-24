package main

import (
	"fmt"
	"math"
)

// float 类型可以作为 map 的 key 吗?
func main() {
	m := make(map[float64]int)
	m[2.4] = 2
	m[3.4] = 3
	m[math.NaN()] = 1
	m[math.NaN()] = 2

	for k, v := range m {
		fmt.Printf("[%v : %d]\n", k, v)
	}

	fmt.Printf("k: %v, v:%d\n", math.NaN(), m[math.NaN()])
	fmt.Printf("k: %v, v:%d\n", 2.4000000000000000000001, m[2.4000000000000000000001]) // 精度丢失
	fmt.Println(math.NaN() == math.NaN())
	// [3.4 : 3]
	// [NaN : 1]
	// [NaN : 2]
	// [2.4 : 2]
	// k: NaN, v:0
	// k: 2.4, v:2
	// false
}
