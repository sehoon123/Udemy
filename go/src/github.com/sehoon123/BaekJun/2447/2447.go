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

// func clear(stars [][]int, n int) {
// 	if n == 1 {
// 		return
// 	}
// 	for i := 0; i < n/3; i++ {
// 		for j := 0; j < n/3; j++ {
// 			stars[n/3+i][n/3+j] = 1
// 		}
// 	}
// 	clear(stars, n/3)
// }

func clear(stars [][]int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i%3 == 1 && j%3 == 1 {
				stars[i][j] = 1
			} else if (i/3)%3 == 1 && (j/3)%3 == 1 {
				stars[i][j] = 1
			} else if (i/9)%3 == 1 && (j/9)%3 == 1 {
				stars[i][j] = 1
			} else if (i/27)%3 == 1 && (j/27)%3 == 1 {
				stars[i][j] = 1
			} else if (i/81)%3 == 1 && (j/81)%3 == 1 {
				stars[i][j] = 1
			} else if (i/243)%3 == 1 && (j/243)%3 == 1 {
				stars[i][j] = 1
			} else if (i/729)%3 == 1 && (j/729)%3 == 1 {
				stars[i][j] = 1
			} else if (i/2187)%3 == 1 && (j/2187)%3 == 1 {
				stars[i][j] = 1
			} else if (i/6561)%3 == 1 && (j/6561)%3 == 1 {
				stars[i][j] = 1
			}
		}
	}
}
