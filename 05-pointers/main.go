package main

import "fmt"

func main() {
	fmt.Println("Learn Pointers")

	var ptr *int
	fmt.Println("The value of pointer", ptr)

	myNumber := 56
	var ptr1 = &myNumber

	fmt.Println("The value of pointer is ", ptr1)
	fmt.Println("The value of pointer is ", *ptr1)

	*ptr1 = *ptr1 + 2
	fmt.Println("The final value is", myNumber)

}
