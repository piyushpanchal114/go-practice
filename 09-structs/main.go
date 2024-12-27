package main

import "fmt"

func main() {
	fmt.Println("Structs in go...")

	user := User{Name: "Piyush", Email: "piyush@piyush.com", Status: true, Age: 21}
	fmt.Println("User details", user)
	fmt.Printf("User with more details %+v\n", user)
	fmt.Printf("The Email of %v is %v\n", user.Name, user.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
