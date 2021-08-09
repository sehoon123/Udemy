package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	w.Flush()

	var N int
	fmt.Fscanf(r, "%d\n", &N)
	var weight int
	var height int
	var list [][]int

	for i := 0; i < N; i++ {
		fmt.Fscanf(r, "%d %d\n", &weight, &height)
		tmp := make([]int, 2)
		tmp[0] = weight
		tmp[1] = height
		list = append(list, tmp)
	}

	var ranking []int

	for i := 0; i < N; i++ {
		var rank int
		for j := 0; j < N; j++ {
			if list[i][0] < list[j][0] && list[i][1] < list[j][1] {
				rank++
			}
		}
		ranking = append(ranking, rank+1)
	}

	for _, v := range ranking {
		fmt.Printf("%d ", v)

	}
}
