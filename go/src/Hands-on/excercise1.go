package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Goroutines :", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("hello from one")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello from two")
		wg.Done()
	}()

	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Goroutines :", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Goroutines :", runtime.NumGoroutine())

	fmt.Println("about to exit")
}
