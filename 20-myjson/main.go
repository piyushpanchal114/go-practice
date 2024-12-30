package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("JSON in Go lang")
	// EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
                "name": "Reactjs",
                "price": 299,
                "channel": "chai aur code",
                "tags": [
                        "js",
                        "frontend"
                ]
        }
	`)
	var newCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &newCourse)
		fmt.Printf("Json is: \n %#v\n", newCourse)
	} else {
		fmt.Println("Invalid JSON")
	}

	// Other cases
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("Json is: \n %#v\n", myOnlineData)

	// Looping the map data
	for k, v := range myOnlineData {
		fmt.Printf("The key is '%v' and the data is '%v' and Type is '%T'\n", k, v, v)
	}

}
