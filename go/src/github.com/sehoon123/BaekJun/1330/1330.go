package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	switch {

	case a > b:
		fmt.Println(">")
	case a == b:
		fmt.Println("==")
	case a < b:
		fmt.Println("<")
	}
}
