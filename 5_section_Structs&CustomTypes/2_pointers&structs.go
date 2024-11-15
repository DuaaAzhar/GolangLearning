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
	displayUserData(&appUser)

}

func getUserData(prompt string) string {
	fmt.Println(prompt)
	var data string
	fmt.Scan(&data)
	return data
}

func displayUserData(u *user) {
	// for structs we dont need to destructure pointer as Go already gives us the shortcut by u.firstName
	// direct way to using pointer structs is
	fmt.Println((*u).firstName, (*u).lastName, (*u).DOB)
	fmt.Println(u.firstName, u.lastName, u.DOB)
}
