package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {

	s := `[{"First":"James","Last":"Bond","Age":23},{"First":"Miss"
	,"Last":"MoenyPenny","Age":32}]`

	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {

		fmt.Println(err)
	}

	fmt.Println("all of the data", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
