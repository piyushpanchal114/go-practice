package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("JSON in Go lang")
	EncodeJson()
}

type course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Platform string   `json:"channel"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {
	cacCourses := []course{
		{"Reactjs", 299, "chai aur code", "abc123", []string{"js", "frontend"}},
		{"Django", 399, "chai aur code", "abc123", []string{"python", "fullstack"}},
		{"Nextjs", 199, "chai aur code", "abc123", []string{"js", "fullstack"}},
		{"Chai Live", 0, "chai aur code", "abc123", nil},
	}

	finalJson, err := json.MarshalIndent(cacCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
