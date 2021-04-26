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

	field := make([][]int, n)

	for i := 0; i < n; i++ {
		field[i] = make([]int, n)
	}

	clear(field, n)

	for i, _ := range field {
		for _, j := range field[i] {
			if j == 0 {
				fmt.Fprintf(w, "*")
			} else {
				fmt.Fprintf(w, " ")
			}
		}
		fmt.Fprintf(w, "\n")
	}

	// star(n)
}

func clear(stars [][]int, n int) {
	if n == 1 {
		return
	}
	for i := 0; i < n/3; i++ {
		for j := 0; j < n/3; j++ {
			stars[n/3+i][n/3+j] = 1
		}
	}
	clear(stars, n/3)
}
