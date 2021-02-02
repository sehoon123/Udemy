package main

import "fmt"

func main() {
	var x, y int16
	fmt.Scan(&x, &y)
	switch {
	case x > 0 && y > 0:
		fmt.Println(1)
	case x > 0 && y < 0:
		fmt.Println(4)
	case x < 0 && y < 0:
		fmt.Println(3)
	case x < 0 && y > 0:
		fmt.Println(2)
	}
}
