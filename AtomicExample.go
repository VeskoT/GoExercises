package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var incrementer int64 = 0
	var wg sync.WaitGroup
	counter := 100
	wg.Add(counter)
	for i:=0; i < counter; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Incrementer value: ", incrementer)
}
