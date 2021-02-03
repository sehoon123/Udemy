package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	str, _ := r.ReadString('\n')
	str = strings.TrimSpace(str)
	n, _ := strconv.Atoi(str)

	star := "*"

	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			w.WriteString(" ")
		}
		w.WriteString(star)
		w.WriteByte('\n')
		star += "*"
	}
	w.Flush()

}
