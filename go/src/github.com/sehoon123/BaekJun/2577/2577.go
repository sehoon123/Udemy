package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var a, b, c int
	fmt.Fscan(r, &a, &b, &c)

	num := strconv.Itoa(a * b * c)

	array := map[int]int{}

	for _, v := range num {
		array[int(v)-'0'] += 1
	}

	for i := 0; i < 10; i++ {
		fmt.Fprintln(w, array[i])
	}
}
