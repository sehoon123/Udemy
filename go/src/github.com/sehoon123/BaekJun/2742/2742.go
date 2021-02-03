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
	N, _ := strconv.Atoi(str)

	for i := N; i >= 1; i-- {

		w.WriteString(strconv.Itoa(i) + "\n")
	}
	w.Flush()

}
