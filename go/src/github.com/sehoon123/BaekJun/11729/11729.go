package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin),bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n int
	fmt.Fscanf(r,"%d\n", &n)
	fmt.Fprintf(w,"%.f\n", math.Pow(2,float64(n))-1)
	hanoi(w, n,1,2,3)

}



func hanoi(w *bufio.Writer,N, start, mid, to int) {
	if N == 1 {
		fmt.Fprintf(w,"%d %d\n", start, to)
		return
	}

	hanoi(w,N-1, start, to, mid)

	fmt.Fprintf(w,"%d %d\n", start, to)

	hanoi(w,N-1, mid, start, to)



}
