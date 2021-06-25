package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[int]int{1: 11, 3: 33, 7: 77, 8: 88, 4: 44, 2: 22, 5: 55, 6: 66, 9: 99}
	fmt.Println("before sort:")
	for k, v := range m {
		fmt.Printf("[%d:%d]\n", k, v)
	}

	fmt.Println("after sort:")
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("[%d:%d]\n", k, m[k])
	}
}
