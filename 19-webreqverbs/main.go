package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb in Go lang")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
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

func PerformPostJsonRequest() {
	const myurl = "https://api.freeapi.app/api/v1/users/register"

	body := strings.NewReader(`
		{
			"email": "user.email@domain.com",
			"password": "test@123",
			"role": "ADMIN",
			"username": "doejohn"
		}
	`)
	response, err := http.Post(myurl, "application/json", body)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	databyte, _ := io.ReadAll(response.Body)

	var content strings.Builder
	content.Write(databyte)
	fmt.Println("content", content.String())

}

func PerformPostFormRequest() {
	const myurl = "https://api.freeapi.app/api/v1/ecommerce/products"

	data := url.Values{}
	data.Add("category", "649865ab297b287175aec1d7")
	data.Add("description", "This is description")
	data.Add("name", "abc")
	data.Add("price", "55")
	data.Add("stock", "10")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	databyte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var content strings.Builder
	content.Write(databyte)
	fmt.Println("resp1", content.String())
}
