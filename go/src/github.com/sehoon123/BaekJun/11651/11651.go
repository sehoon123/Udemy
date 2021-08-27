package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var scanner = bufio.NewScanner(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func nextInt() int {
	scanner.Scan()
	r, f := 0, 1
	for _, c := range scanner.Bytes() {
		if c == '-' {
			f = -1
			continue
		}
		r *= 10
		r += int(c - '0')
	}
	return r * f
}

func main() {
	scanner.Split(bufio.ScanWords)
	n := nextInt()
	a := make([][2]int, n)

	for i := range a {
		a[i][0], a[i][1] = nextInt(), nextInt()
	}

	sort.Slice(a, func(i, j int) bool {
		if a[i][1] != a[j][1] {
			return a[i][1] < a[j][1]
		}
		return a[i][0] < a[j][0]
	})

	for _, v := range a {
		fmt.Fprintln(writer, v[0], v[1])
	}
	defer writer.Flush()
}
