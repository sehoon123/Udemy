package main

import "fmt"

func main() {
	//map[key]value
	m := map[string]int{
		"James":           32,
		"Miss Moenypenny": 27,
	}
	fmt.Println(m)
	delete(m, "James")
	fmt.Println(m)

	delete(m, "Ian Fleming")
	fmt.Println(m)
	fmt.Println(m["Miss Moenypenny"])
	fmt.Println(m["Ian Felming"])

	if v, ok := m["Miss Moenypenny"]; ok {
		fmt.Println("value", v)
	}
}
