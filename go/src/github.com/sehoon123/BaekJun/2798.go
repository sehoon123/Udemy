package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	var m int
	fmt.Scanf("%d %d\n", &n, &m)

	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	str, _ := r.ReadString('\n')
	strs := strings.Fields(str)
	nums := []int{}

	for _, i := range strs {
		j, _ := strconv.Atoi(i)

		nums = append(nums, j)
	}

	var result int
	var max int

	for i, v1 := range nums[:n-3] {
		for _, v2 := range nums[i+1:n-2]{
			for _, v3 := range nums[i+2:n-1] {
				result = v1 + v2 + v3
				if result > max {
					max = result
				}
			}
		}
	}

	fmt.Println(max)






}
