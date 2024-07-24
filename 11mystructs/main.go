package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	Arijit := User{"Arijit", "arijitdas.hashnode.dev", true, 19}
	fmt.Println(Arijit)
	fmt.Printf("Arijit details are: %+v\n", Arijit)
	fmt.Printf("Name is %v and email is %v.", Arijit.Name, Arijit.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
