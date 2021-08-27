package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()

	for _, c := range s.Bytes() {
		fmt.Println(c - '0')
	}
}
