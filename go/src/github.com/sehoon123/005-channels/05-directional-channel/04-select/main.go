package main

import (
	"fmt"
	"sync"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)

	//get value channel
	fanin := make(chan int)

	//send
	go send(eve, odd)

	// receive
	go receive(eve, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func receive(eve, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range eve {
			fanin <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}

func send(eve, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(eve)
	close(odd)
}
