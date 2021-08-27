package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	s, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	s.Scan()

	n, _ := strconv.Atoi(s.Text())

	numList := make([]int, 8001)

	for i := 0; i < n; i++ {
		s.Scan()
		num, _ := strconv.Atoi(s.Text())
		numList[num+4000]++
	}

	var sortedList []int

	for i, v := range numList {
		for j := 0; j < v; j++ {
			sortedList = append(sortedList, i-4000)
		}
	}
	fmt.Fprintf(w, "%v\n", mean(sortedList))
	fmt.Fprintf(w, "%v\n", median(sortedList))
	fmt.Fprintf(w, "%v\n", Mode(numList))
	fmt.Fprintf(w, "%v\n", scope(sortedList))

}

func mean(array []int) float64 {
	var sum int
	for _, v := range array {
		sum += v
	}

	return math.Round(float64(sum) / float64(len(array)))
}

func median(array []int) int {
	return array[len(array)/2]
}

func Mode(array []int) int {
	var flag int
	var max int
	var first int
	var second int

	for i, v := range array {
		if v > max {
			max = v
			first = i - 4000
		}
	}

	for i, v := range array {
		if v == max {
			flag++
		}
		if flag == 2 {
			second = i - 4000
			return second
		}
	}
	return first
}

func scope(array []int) int {
	return array[len(array)-1] - array[0]
}
