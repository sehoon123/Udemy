package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	str, _ := r.ReadString('\n')
	strs := strings.Fields(str)
	// n1, _ := strconv.Atoi(strs[0])

	str, _ = r.ReadString('\n')
	strs = strings.Fields(str)

	var sum int
	for i := range strs[0] {
		num, _ := strconv.Atoi(string(strs[0][i]))
		sum += num
	}

	fmt.Fprintln(w, sum)

}
