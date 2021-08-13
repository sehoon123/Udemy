package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	defer w.Flush()

	fmt.Fscanf(r, "%d\n", &n)
	numList := make([]int, 10001)

	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		numList[num]++
	}

	for i, v := range numList {
		for j := 0; j < v; j++ {
			fmt.Fprintf(w, "%d\n", i)
		}
	}

}
