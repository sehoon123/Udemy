package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(len(a))
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
	fmt.Println(a[3])
	fmt.Println(a[4])

	for i, v := range a {
		fmt.Println(i, v)
	}
}
