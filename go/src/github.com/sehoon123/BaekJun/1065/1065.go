package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	fmt.Fprint(w, Han(n))

}

var wg sync.WaitGroup

func Han(n int) int {
	var count int
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < n+1; i++ {
			if i < 100 {
				count++
			} else if (i/100 + i%10) == (i/10-(i/100)*10)*2 {
				count++
			}
		}
	}()
	wg.Wait()
	return count
}
