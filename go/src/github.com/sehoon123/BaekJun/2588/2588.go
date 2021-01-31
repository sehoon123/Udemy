package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	n2 := strconv.Itoa(b)
	b1, _ := strconv.Atoi(string(n2[2]))
	b2, _ := strconv.Atoi(string(n2[1]))
	b3, _ := strconv.Atoi(string(n2[0]))
	fmt.Println(a * b1)
	fmt.Println(a * b2)
	fmt.Println(a * b3)
	fmt.Println(a * b)
}
