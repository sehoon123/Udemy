package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Axis struct {
	X int
	Y int
}

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	a := make([]Axis, n)
	for i := 0; i < n; i++ {
		var ta Axis
		s, _ := r.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		split := strings.Split(s, " ")
		ta.X, _ = strconv.Atoi(split[0])
		ta.Y, _ = strconv.Atoi(split[1])
		a[i] = ta
	}
	sort.SliceStable(a, func(i, j int) bool {
		return a[i].Y < a[j].Y
	})
	sort.SliceStable(a, func(i, j int) bool {
		return a[i].X < a[j].X
	})

	for _, v := range a {
		fmt.Fprintln(w, v.X, v.Y)
	}

}
