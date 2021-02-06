package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	numlist := []int{}

	var x int

	for i := 0; i < 10; i++ {
		fmt.Fscan(r, &x)
		numlist = append(numlist, x)
	}

	remainder := map[int]int{}

	for _, v := range numlist {
		remainder[v%42] += 1
	}

	fmt.Fprint(w, len(remainder))
}
