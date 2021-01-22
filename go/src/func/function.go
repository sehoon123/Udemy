package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scanf("%f", &x)
	a := (((math.Sqrt(((4 * x) - 1) / 3)) - 1) / 2) + 1
	if math.Mod(a, 1) == 0 {
		fmt.Println(int(a))
	} else {
		fmt.Println(int(a) + 1)
	}

}
