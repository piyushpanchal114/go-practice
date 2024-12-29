package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://piyushpanchal.com"
const url2 = "https://google.com"

func main() {
	fmt.Println("Handling Web Request in Go lang")

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// fmt.Println("the resp is", resp)
	fmt.Printf("The response is of type: %T\n", resp)
	defer resp.Body.Close()

	databyte, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Println("content", content)
}
