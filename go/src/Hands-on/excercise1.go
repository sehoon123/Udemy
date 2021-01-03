package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken", "not stirred", "Martinis", "Women"},
		"moenypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
	}

	for k, v := range m {
		fmt.Printf("Key : %v\t, Value : %v\n", k, v)
		fmt.Println("\t What I like : ")
		for _, like := range v {
			fmt.Printf("\t\t %v\n", like)
		}
	}

}
