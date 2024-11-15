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

	// STRUCT LITERAL

	// var appUser user
	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	DOB:       userBirthDate,
	// 	createdAt: time.Now(),
	// }

	// OR   //Recommended

	var appUser user = user{
		firstName: userFirstName,
		lastName:  userLastName,
		DOB:       userBirthDate,
		createdAt: time.Now(),
	}
	displayUserData(appUser)

	// OR
	// var appUser user = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthDate,
	// 	time.Now(),
	// }

	// output user details

}

func getUserData(prompt string) string {
	fmt.Println(prompt)
	var data string
	fmt.Scan(&data)
	return data
}

func displayUserData(u user) {
	fmt.Println(u.firstName, u.lastName, u.DOB)
}
