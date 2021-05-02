package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int

	fmt.Fscanf(r, "%d\n", &n)
	fmt.Fprintf(w, "%.f\n", math.Pow(2, float64(n))-1)
}

func hanoi(k int) {
	if k%2 == 0 {
	}
}
