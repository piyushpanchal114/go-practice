package main

import "fmt"

func main() {
	fmt.Println("Functions in Go")
	greeter()
	fmt.Println("The adder is", adder(5, 9))
	fmt.Println("The adderPro", adderPro(1, 2, 3, 4, 5))
}

func greeter() {
	fmt.Println("Namaste from Go")
}

func adder(x int, y int) int {
	return x + y
}

func adderPro(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}
