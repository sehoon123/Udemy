package main

import "fmt"

func main() {
	var year int16
	fmt.Scan(&year)
	switch {
	case year%400 == 0:
		fmt.Println(1)
	case year%100 == 0:
		fmt.Println(0)
	case year%4 == 0:
		fmt.Println(1)
	default:
		fmt.Println(0)
	}
}
