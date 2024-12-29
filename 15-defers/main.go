package main

import "fmt"

func main() {
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("World-")
	fmt.Println("Defers in Go")
	fmt.Println("Hello")
	defer looper()
}

func looper() {
	for i := 1; i < 5; i++ {
		defer fmt.Print(i)
	}
}
