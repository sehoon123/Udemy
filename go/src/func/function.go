package main

import "fmt"

func main() {
	foo()
	bar("Sehoon")
}

func foo() {
	fmt.Println("Hello, from foo")
}

// everything is Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello", s)
}
