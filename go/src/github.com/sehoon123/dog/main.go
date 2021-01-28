package main

import (
	"/home/ghostb123/Udemy/go/src/dog/dog"
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
