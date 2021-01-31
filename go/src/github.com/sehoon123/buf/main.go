package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("/home/ghostb123/Udemy/go/src/github.com/sehoon123/buf/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text(), lineCount)
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(lineCount)
}
