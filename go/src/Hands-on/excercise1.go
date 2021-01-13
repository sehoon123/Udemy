package main

import "fmt"

func main() {
	a := foo()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := foo()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
