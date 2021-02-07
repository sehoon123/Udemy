package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	list := make(map[rune]rune, 26)

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	str, _ := r.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	str = strings.ToUpper(str)

	for _, v := range str {
		list[(v)-65] += 1
	}

	var maxNumber rune
	var x rune
	var count int

	for n := range list {
		if list[n] > maxNumber {
			maxNumber = list[n]
			x = rune(n)
		}

	}

	for n := range list {
		if list[n] == maxNumber {
			count++
		}

	}

	if count != 1 {
		fmt.Fprint(w, "?")
	} else {
		fmt.Fprintf(w, "%s", string(x+65))
	}
}
