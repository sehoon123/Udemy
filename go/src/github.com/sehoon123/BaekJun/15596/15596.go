package main

func main() {
	var a []int
	sum(a)
}

func sum(a []int) int {
	var sum int
	for _, v := range a {
		sum += v
	}
	return sum
}
