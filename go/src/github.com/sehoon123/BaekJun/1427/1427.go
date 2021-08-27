package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	s.Scan()
	n := s.Text()
	numList := make([]int, 10)

	for _, v := range n {
		tmp, _ := strconv.Atoi(string(v))
		numList[tmp]++
	}

	for i := 9; i >= 0; i-- {
		for j := 0; j < numList[i]; j++ {
			fmt.Fprintf(w, "%d", i)
		}
	}
}
