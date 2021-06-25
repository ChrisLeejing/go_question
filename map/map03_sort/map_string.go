package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"a": 1, "d": 4, "c": 3, "f": 6, "e": 5, "b": 2}
	fmt.Println("before sort: ")
	for k, v := range m {
		fmt.Printf("[%s:%d]\n", k, v)
	}
	fmt.Println("after sort: ")

	var keys []string
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("[%s:%d]\n", k, m[k])
	}
}
