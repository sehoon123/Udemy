package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c, q <-chan int) {
	for {

		select {
		case <-c:
			fmt.Println("receive from c", <-c)
		case <-q:
			fmt.Println("receive from q", <-q)
			return
		}
	}

}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {

		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
		close(q)
	}()

	return c
}
