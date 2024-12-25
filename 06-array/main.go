package main

import "fmt"

func main() {
	var fruitList [4]string

	fruitList[0] = "Mango"
	fruitList[1] = "Apple"
	fruitList[3] = "Banana"

	fmt.Println("Here is the fruit list", fruitList)
	fmt.Println("Length of the fruit list", len(fruitList))

	var vegList = [5]string{"Potato", "Tomato", "Cabbage"}
	fmt.Println("This is the veggy list", vegList)
	fmt.Println("Length of veg list", len(vegList))
}
