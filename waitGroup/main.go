package main

import (
	"sync"
)
d
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	add(12, 114)
	go add(22, 33)
	wg.Wait()

}

func add(a int, b int) int {
	defer wg.Done()
	c := a + b
	return c

}
