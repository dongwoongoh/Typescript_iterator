package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sum(a, b int) {
	s := 0
	for i := a; i <= b; i++ {
		s += i
	}
	fmt.Println(s)
	wg.Done()
}

func main() {
	try := 10
	wg.Add(try)
	for i := 0; i < 10; i++ {
		go sum(1, 100)
	}
	wg.Wait()
	fmt.Println("multiple thread has been done")
}
