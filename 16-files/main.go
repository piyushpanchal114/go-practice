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
	raiseErr(err)

	length, err := io.WriteString(file, content)

	raiseErr(err)
	fmt.Println("length", length)
	defer file.Close()
	readFile("./temp.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	raiseErr(err)
	fmt.Println("datavyte", databyte)
	fmt.Println("string", string(databyte))
}

func raiseErr(err error) {
	if err != nil {
		panic(err)
	}
}
