package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(100)

	counter := 0

	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println(runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(counter)

}
