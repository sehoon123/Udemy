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
	list := make([]int, 26)

	for i := 0; i < 26; i++ {
		list[i] = -1
	}

	for i, v := range str {
		if list[v-97] == -1 {
			list[v-97] = (i)
		}
	}

	for _, v := range list {
		fmt.Fprintf(w, "%d ", v)
	}

}
