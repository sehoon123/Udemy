package main

import (
	"fmt"
	"math"
	"time"
)

func fiboarray(n float64) float64 {
	return 1/math.Sqrt(5)*(math.Pow(1+math.Sqrt(5), n))/2 - (math.Pow(1-math.Sqrt(5), n))/2
}

func fib(n int64, cache map[int64]int64) int64 {
	if n == 1 || n == 2 {
		return 1
	}
	if cache[n] == 0 {
		cache[n] = fib(n-1, cache) + fib(n-2, cache)
	}
	return cache[n]
}

func main() {
	startTime := time.Now()

	fiboarray(140)

	e := time.Since(startTime)
	fmt.Printf("%v", e)
}
