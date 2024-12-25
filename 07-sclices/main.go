package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Go... Sclices")

	var fruitSclices = []string{"Apple", "Banana", "Peach"}
	fruitSclices = append(fruitSclices, "Mango", "Papaya")
	fmt.Println("fruitsciles", fruitSclices)

	fruitSclices = append(fruitSclices[1:])
	fmt.Println("fruitSclices", fruitSclices)

	highScores := make([]int, 4)

	highScores[0] = 265
	highScores[1] = 565
	highScores[2] = 295
	highScores[3] = 789
	// highScores[4] = 988

	highScores = append(highScores, 564, 941, 326)
	fmt.Println("hign scores", highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println("sorted list", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// remove elements from sclices using append method

	var courses = []string{"reactjs", "python", "swift", "java", "kotlin"}
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("courses", courses)
}
