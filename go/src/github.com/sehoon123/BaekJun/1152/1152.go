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
	strs := strings.Fields(str)

	var count int
	for range strs {
		count++
	}

	fmt.Fprintf(w, "%d\n", count)

}
