package main

import "fmt"

func main() {
	fmt.Println("Maps in Go lang...")

	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println("Total languages", languages)
	fmt.Println("js is short form of", languages["js"])

	delete(languages, "rb")
	fmt.Println("Total languages", languages)

	// for loop in go lang

	for key, value := range languages {
		fmt.Printf("For key %v, Value is %v\n", key, value)
	}

}
