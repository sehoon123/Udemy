package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n, m int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanf(r, "%d %d\n", &n, &m)

	var mat = make([]string, n)
	//한줄씩 받아오기
	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%v\n", &mat[i])

	}
	// 초기에 다 바꿔야 되는 상황을 min으로 설정
	var min = n * m

	for i := 0; i < n-7; i++ {
		for j := 0; j < m-7; j++ {
			var c1 = float64(0)
			var c2 = float64(0)
			for k := i; k < i+8; k++ {
				for l := j; l < j+8; l++ {
					if (k+l)%2 == 0 {
						if string(mat[k][l]) == "B" {
							c1++
						} else {
							c2++
						}
					} else {
						if string(mat[k][l]) == "B" {
							c2++
						} else {
							c1++
						}
					}
				}
			}
			if min > int(math.Min(c1, c2)) {
				min = int(math.Min(c1, c2))
			}
		}
	}
	fmt.Println(min)
}
