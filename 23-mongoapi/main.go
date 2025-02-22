package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/piyushpanchal114/mongoapi/routers"
)

func main() {
	fmt.Println("Mongo API")
	r := routers.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000...")
}
