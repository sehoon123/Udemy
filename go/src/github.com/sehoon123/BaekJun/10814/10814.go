package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscanf(r, "%d\n", &n)

	mans := make([][2]string, n)

	for i := 0; i < n; i++ {
		var name string
		var age string

		fmt.Fscanf(r, "%v %v\n", &age, &name)
		mans[i][0], mans[i][1] = age, name
	}

	sort.Slice(mans, func(i, j int) bool {
		return mans[i][0] < mans[j][0]
	})

	for _, l1 := range mans {
		fmt.Fprintf(w, "%s %s\n", l1[0], l1[1])
	}

}
