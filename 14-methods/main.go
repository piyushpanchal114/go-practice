package main

import "fmt"

func main() {
	fmt.Println("Methods in go...")

	user := User{Name: "Piyush", Email: "piyush@piyush.com", Status: true, Age: 21}
	fmt.Println("User details", user)
	fmt.Printf("User with more details %+v\n", user)
	fmt.Printf("The Email of %v is %v\n", user.Name, user.Email)

	user.GetStatus()
	user.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Printf("The status of %v is %v\n", u.Name, u.Status)
}

func (u User) NewMail() {
	u.Email = "temp@piyush.com"
	fmt.Printf("The new mail for %v is %v", u.Name, u.Email)
}
