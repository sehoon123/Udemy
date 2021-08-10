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
	for k := 0; k < n; k++ {
		for i := 0; i < n-1; i++ {
			var tmp int

			if list[i] > list[i+1] {
				tmp = list[i]
				list[i] = list[i+1]
				list[i+1] = tmp
			}
		}
	}

	for _, v := range list {
		fmt.Println(v)
	}
}
