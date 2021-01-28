package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("vim-go")
	fmt.Fprintln(os.Stdout, "Hello, playground")
	io.WriteString(os.Stdout, "Hello, playground")
}
