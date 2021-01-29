package main

import (
	"fmt"

	"github.com/sehoon123/Hands-on/13-2/quote"
	"github.com/sehoon123/Hands-on/13-2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
