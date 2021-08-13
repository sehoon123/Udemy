package main

import (
	"bufio"
	"fmt"
	"os"
)

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func main() {
	var n int
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscanf(r, "%d\n", &n)

	list := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%d\n", &list[i])
	}

	sorted := mergeSort(list)

	for i := 0; i < n; i++ {
		fmt.Fprintf(w, "%d\n", sorted[i])
	}
}
