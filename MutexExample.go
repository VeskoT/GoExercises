package main

import (
	"fmt"
	"sync"
)

func main() {

	var incrementer int = 0
	var wg sync.WaitGroup
	counter := 100
	wg.Add(counter)
	var mutex = sync.Mutex{}
	for i:=0; i < counter; i++ {
		go func() {
			mutex.Lock()
			value := incrementer
			value++
			incrementer = value
			fmt.Println(incrementer)
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Incrementer value: ", incrementer)
}
