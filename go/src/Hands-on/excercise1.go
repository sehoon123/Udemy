package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person, n int) {
	p.age = n
}

func main() {
	p1 := person{
		first: "Sehun",
		last:  "Jung",
		age:   23,
	}
	fmt.Println(p1)

	changeMe(&p1, 10)

	fmt.Println(p1)
}
