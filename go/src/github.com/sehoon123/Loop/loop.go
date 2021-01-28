package main

import "fmt"

func main() {
	n := "Bond"
	switch n {
	case "jack", "Bond", "whatever":
		fmt.Println("I'm bond")
	case "mom":
		fmt.Println("I'm your mother")
	case "sehoon":
		fmt.Println("I'm James Bond")
	default:
		fmt.Println("this is default")
	}
}
