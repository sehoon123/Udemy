package main

import "fmt"

func main() {
	c := make(chan int)
	add(c)
	for v := range c {
		fmt.Println(v)

	}

}

func add(c chan int) {
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
}
