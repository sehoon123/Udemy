package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int)
	cs := make(chan<- int)

	cr <- 42
	cr <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("==============")
	fmt.Printf("%T\n", c)

}
