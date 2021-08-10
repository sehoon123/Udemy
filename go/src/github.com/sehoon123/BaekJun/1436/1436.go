package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	count := 0

	for i := 666; i < 10000666; i++ {
		if strings.Contains(strconv.Itoa(i), "666") {
			count++
		}
		if count == n {
			fmt.Println(i)
			break
		}
	}

}
