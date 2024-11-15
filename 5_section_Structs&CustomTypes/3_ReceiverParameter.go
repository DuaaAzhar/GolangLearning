package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	DOB       string
	createdAt time.Time
}

// Defining Method having receiver param of type user struct
func (u user) displayUserData() {
	fmt.Println(u.firstName, u.lastName, u.DOB)
}

// Creating a Receiver param method that mutates the struct value
// if we pass the value like this (u user) it will not mutate struct bcz receiver methods pass only copies not original value
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := getUserData("Enter your First Name: ")
	userLastName := getUserData("Enter your Last Name: ")
	userBirthDate := getUserData("Enter your BirthDate (MM/DD/YYYY): ")

	var appUser user = user{
		firstName: userFirstName,
		lastName:  userLastName,
		DOB:       userBirthDate,
		createdAt: time.Now(),
	}

	// To call the struct method, u dont need to pass any param as Go will automatically pass the copy of user struct to method
	appUser.displayUserData()
	appUser.clearUserName()
	appUser.displayUserData()

}

func getUserData(prompt string) string {
	fmt.Println(prompt)
	var data string
	fmt.Scan(&data)
	return data
}
