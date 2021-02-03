package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	strNum, _ := r.ReadString('\n')
	Nums := strings.Fields(strNum)
	n1, _ := strconv.Atoi(Nums[0])
	n2, _ := strconv.Atoi(Nums[1])

	list := []int{}

	var a int
	for i := 0; i < n1; i++ {
		fmt.Scanf("%d", &a)
		list = append(list, a)

	}

	var ans string

	for _, v := range list {
		if v < n2 {
			ans += strconv.Itoa(v) + " "
		}

	}
	answer := strings.TrimRight(ans, " ")
	w.WriteString(answer)
	w.Flush()
}
