package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var x int
	numlist := make([]int, 9)

	for i := 0; i < 9; i++ {
		fmt.Fscan(r, &x)
		numlist[i] = x
	}

	var max, index int
	max = numlist[0]

	for i, v := range numlist {
		if v >= max {
			max = v
			index = i + 1
		}
	}
	fmt.Fprintf(w, "%d\n%d", max, index)
}
