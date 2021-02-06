package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	list := []string{}
	var temp string
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &temp)
		list = append(list, temp)
	}

	test := make([][]int, n)

	var count int
	for i, v := range list {
		count = 0
		for _, k := range v {
			if string(k) == "O" {
				count += 1
			} else {
				test[i] = append(test[i], count)
				count = 0
			}
		}
		test[i] = append(test[i], count)

	}

	sum := 0
	for _, v := range test {
		sum = 0
		for _, k := range v {
			sum += k * (k + 1) / 2
		}
		fmt.Fprintf(w, "%d\n", sum)
	}

}
