package main

import (
	"fmt"
	"time"
)

// Define a struct
type user struct {
	firstName string
	lastName string 
	birthDate string
	createdAt time.Time
}

// Define and attach method to struct 
func (u user) outputUserDetail() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user
	appUser = user{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}
	
	appUser.outputUserDetail()
}



func getUserData(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scan(&value)
	return value
}