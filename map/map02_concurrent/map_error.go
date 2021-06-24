package main

import "fmt"

func main() {
	m := make(map[int]int)

	for i := 0; i < 100; i++ {
		go func() {
			m[i] = i
		}()
	}

	fmt.Println("len(m) = ", len(m))
}
