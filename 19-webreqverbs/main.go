package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb in Go lang")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "https://api.freeapi.app/api/v1/public/stocks"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println("status code", response.StatusCode)
	fmt.Println("content length", response.ContentLength)
	defer response.Body.Close()

	// One way
	databyte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("data", string(databyte))

	// Other way which is more preffered way
	var responseBuilder strings.Builder
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	byteCount, _ := responseBuilder.Write(content)
	fmt.Println("Byte count", byteCount)
	fmt.Println("the data is:", responseBuilder.String())

}
