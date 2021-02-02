package main

import "fmt"

func main() {
	var h, m int8
	fmt.Scan(&h, &m)
	switch {
	case m >= 45:
		fmt.Printf("%d %d", h, m-45)
	case h >= 1 && m < 45:
		fmt.Printf("%d %d", h-1, m-45+60)
	case h == 0 && m < 45:
		fmt.Printf("%d %d", 23, m-45+60)
	}
}
