package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n, x int

	fmt.Fscanln(r, &n, &x)
	defer w.Flush()

	var list = make([]int, n)
	for i := range list {
		fmt.Fscanf(r, "%d ", &list[i])
		if list[i] < x {
			fmt.Fprintf(w, "%d ", list[i])
		}
	}
	fmt.Fprintf(w, "\n")

}
