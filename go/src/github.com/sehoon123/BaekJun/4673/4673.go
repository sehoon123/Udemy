package main

import (
	"bytes"
	"fmt"
	"strconv"
	"sync"
)

var arr [10001]bool

func d(n int) int {
	sum := n
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10000; i++ {
			result := d(i)
			if result <= 10000 {
				arr[result] = true
			}
		}
	}()

	var buf bytes.Buffer

	arr[0] = true

	for i := range arr {
		if !arr[i] {
			buf.WriteString(strconv.Itoa(i))
			buf.WriteRune('\n')
		}
	}
	fmt.Print(buf.String())
	wg.Wait()

}
