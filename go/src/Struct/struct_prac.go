package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type scagent struct {
	person
	ltk bool
}

func main() {
	sc1 := scagent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   33,
		},
		ltk: true,
	}
	p2 := person{
		first: "Miss",
		last:  "MoneyPenny",
		age:   25,
	}
	fmt.Println(sc1)
	fmt.Println(p2)

	fmt.Println(sc1.first, sc1.last, sc1.age, sc1.ltk)
	fmt.Println(p2.first, p2.last, p2.age)
}
