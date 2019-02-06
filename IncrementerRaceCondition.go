package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var incrementer int = 0
	var wg sync.WaitGroup
	counter := 100
	wg.Add(counter)
	for i:=0; i < counter; i++ {
		go func() {
			value := incrementer
			runtime.Gosched()
			value++
			incrementer = value
			fmt.Println(incrementer)
			wg.Done()
		}()
	}
	fmt.Println("Incrementer value: ", incrementer)
	wg.Wait()
}
