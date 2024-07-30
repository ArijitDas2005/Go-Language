package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	Arijit := User{"Arijit", "Arijit@go.dev", true, 16}
	fmt.Println(Arijit)
	fmt.Printf("Arijit details are: %+v\n", Arijit)
	fmt.Printf("Name is %v and email is %v.\n", Arijit.Name, Arijit.Email)
	Arijit.GetStatus()
	Arijit.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", Arijit.Name, Arijit.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
