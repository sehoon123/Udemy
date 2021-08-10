package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	var count int
	list := make([]int, n)
	for count < n {
		var a int
		fmt.Scanf("%d\n", &a)
		list[count] = a
		count++
	}

	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			var tmp int
			if list[j] < list[j-1] {
				tmp = list[j]
				list[j] = list[j-1]
				list[j-1] = tmp
			}
		}
	}

	fmt.Println(list)
}
