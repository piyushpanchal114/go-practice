package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input:"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our service:")

	// comma ok syntax || try and catch in Go

	input, err := reader.ReadString('\n')
	fmt.Println("Thank you for rating...", input)
	fmt.Printf("Type of rating... %T \n", input)
	fmt.Println("Error.........", err)

}
