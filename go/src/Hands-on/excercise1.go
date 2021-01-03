package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken", "not stirred", "Martinis", "Women"},
		"moenypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
	}

	m["sehoon"] = []string{"I", "Like", "You"}

	for k, v := range m {
		fmt.Println(k, v)
		fmt.Println("\t What I like : ")
		for _, like := range v {
			fmt.Printf("\t\t %v\n", like)
		}
	}

	delete(m, "sehoon")
	fmt.Println(m)

}
