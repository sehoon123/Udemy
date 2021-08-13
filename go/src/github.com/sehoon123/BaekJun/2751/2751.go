package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscanf(r, "%d\n", &n)

	list := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%d\n", &list[i])
	}

	sort.Ints(list)

	for i := 0; i < n; i++ {
		fmt.Fprintf(w, "%d\n", list[i])
	}

}
