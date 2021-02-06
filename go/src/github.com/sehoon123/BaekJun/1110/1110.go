package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var count int
	var n int

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscanln(r, &n)
	count = 1
	temp := (n%10)*10 + ((n/10)+(n%10))%10

	for temp != n {
		count++
		if temp == n {
			break
		}

		temp = (temp%10)*10 + ((temp/10)+(temp%10))%10

	}
	w.WriteString(strconv.Itoa(count))

}
