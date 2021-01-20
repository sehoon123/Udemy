package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Println("My name is ", p.name, "And My age is ", p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "sehoon",
		age:  123,
	}

	p1.speak()

	fmt.Printf("%T\n", p1)

	saySomething(&p1)
	// saySomething(p1)
}
