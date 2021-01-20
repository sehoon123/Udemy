package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
