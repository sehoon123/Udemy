package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	str, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	strs := strings.Fields(str)
	n1, _ := strconv.Atoi(strs[0])
	n2, _ := strconv.Atoi(strs[1])
	fmt.Println(n1 * n2)
}
