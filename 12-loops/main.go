package main

import (
	"fmt"
)

func main() {
	fmt.Println("loops in go lang")
	days := []string{"monday", "tuesday", "wednesday", "thrusday", "friday"}
	// example 1
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// example 2
	fmt.Println("Result of example 2")
	for i := range days {
		fmt.Println(days[i])
	}

	// example 3
	fmt.Println("Result of example 3")
	for i, v := range days {
		fmt.Printf("the index %v has %v value.\n", i, v)
	}

	// example 4 or while loop in go
	fmt.Println("Result of example 4")
	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 3 {
			rougueValue++
			continue
		}
		if rougueValue == 5 {
			goto hit
		}
		fmt.Println("rougue value", rougueValue)

		rougueValue++
	}
hit:
	fmt.Println("THis is goto")

}
