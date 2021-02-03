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

	for i := 1; i <= N; i++ {

		w.WriteString(strconv.Itoa(i) + "\n")
	}
	w.Flush()

}
