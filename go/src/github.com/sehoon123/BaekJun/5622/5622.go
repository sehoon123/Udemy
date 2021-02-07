package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	str, _ := r.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")

	list := []int{}
	for _, v := range str {
		switch {
		case v < 68:
			list = append(list, 2)
		case v < 71:
			list = append(list, 3)
		case v < 74:
			list = append(list, 4)
		case v < 77:
			list = append(list, 5)
		case v < 80:
			list = append(list, 6)
		case v < 84:
			list = append(list, 7)
		case v < 87:
			list = append(list, 8)
		case v < 91:
			list = append(list, 9)
		}
	}

	var sum int
	for _, v := range list {
		sum += v + 1
	}

	fmt.Fprintf(w, "%d\n", sum)
}
