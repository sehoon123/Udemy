package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	s, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	words := make([]string, n)

	for i := range words {
		s.Scan()
		words[i] = s.Text()
	}

	sort.Strings(words)

	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return words[i] < words[j]
		}
		return len(words[i]) < len(words[j])
	})

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if words[i] == words[j] {
				words[j] = ""
			}
		}
	}

	for i, _ := range words {
		if words[i] == "" {
			continue
		}
		fmt.Fprintln(w, words[i])
	}

}
