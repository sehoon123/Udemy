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

	var n int
	fmt.Fscanf(r, "%d\n", &n)

	var nu int

	for i := 0; i < n; i++ {
		var code string
		fmt.Fscanf(r, "%d %s\n", &nu, &code)

		for _, v := range code {
			for i := 0; i < nu; i++ {
				fmt.Fprintf(w, "%c", v)
			}
		}
		fmt.Fprint(w, "\n")
	}

}
