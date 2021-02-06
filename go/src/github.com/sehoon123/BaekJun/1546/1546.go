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
	score := []int{}

	var x int
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &x)
		score = append(score, x)
	}

	max := 0
	for _, v := range score {
		if v > max {
			max = v
		}
	}

	var sum float64
	for _, v := range score {
		sum += (float64(v) / float64(max)) * 100
	}

	fmt.Fprint(w, sum/float64(n))
}
