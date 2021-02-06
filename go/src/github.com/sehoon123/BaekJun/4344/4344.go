package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	var num_case int
	var x int

	test_list := make([][]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &num_case)
		for j := 0; j < num_case; j++ {
			fmt.Fscan(r, &x)
			test_list[i] = append(test_list[i], x)
		}

	}

	for i := 0; i < n; i++ {
		percentage(test_list[i], w)
	}

}

func percentage(arr []int, w io.Writer) {
	var sum float64
	var avg float64
	var count float64
	for _, v := range arr {
		(sum) += float64(v)
	}
	avg = sum / float64(len(arr))

	for _, v := range arr {
		if float64(v) > avg {
			count++
		}
	}
	fmt.Fprintf(w, "%.3f%%\n", count/float64(len(arr))*100)

}
