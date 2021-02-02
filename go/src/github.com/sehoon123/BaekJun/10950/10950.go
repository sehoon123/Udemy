package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}
