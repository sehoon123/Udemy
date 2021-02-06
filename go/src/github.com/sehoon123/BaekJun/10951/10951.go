package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var a, b int

	for {
		_, err := fmt.Scanf("%d %d", &a, &b)
		if err != nil {
			break
		}
		fmt.Fprintln(w, a+b)
	}
}
