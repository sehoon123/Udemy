package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, x int
	fmt.Fscan(r, &n)
	numList := []int{}

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &x)
		numList = append(numList, x)
	}

	var min, max int
	min = numList[0]
	max = numList[0]
	for _, v := range numList {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	w.WriteString(strconv.Itoa(min) + " " + strconv.Itoa(max))

}
