package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var a, b int

	for {
		fmt.Fscanln(r, &a, &b)
		if a == 0 && b == 0 {
			break
		}
		fmt.Fprintf(w, "%d\n", a+b)
	}
}
