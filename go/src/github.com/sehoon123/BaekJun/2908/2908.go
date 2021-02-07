package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var x, y int
	fmt.Fscan(r, &x, &y)
	x = (x%10)*100 + (x % 100) - (x % 10) + x/100
	y = (y%10)*100 + (y % 100) - (y % 10) + y/100

	if x > y {
		fmt.Fprintf(w, "%d\n", x)
	} else {
		fmt.Fprintf(w, "%d\n", y)
	}
}
