package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Handling files in Go")
	content := "This is my First File using Go lang"

	file, err := os.Create("./temp.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("length", length)
	defer file.Close()
	readFile("./temp.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("datavyte", databyte)
	fmt.Println("string", string(databyte))
}
