package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Enter your First Name: ")
	userLastName := getUserData("Enter your Last Name: ")
	userBirthDate := getUserData("Enter your BirthDate (MM/DD/YYYY): ")

	// if we creating struct using pointer
	var appUser *user.User
	appUser, err := user.New(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Println(err)
		return
	}

	// To call the struct method, u dont need to pass any param as Go will automatically pass the copy of user struct to method
	appUser.DisplayUserData()
	appUser.ClearUserName()
	appUser.DisplayUserData()

	// creating embedded Admin struct
	// var admin user.Admin

	admin := user.NewAdmin("duaaazhar@gmail.com", "1234", "my", "admin", "2/3/2000")
	// if in case of defining embedded struct by defining variable
	// admin.U.DisplayUserData()

	// in case of defining direct embedded struct
	admin.DisplayUserData()
	admin.ClearUserName()
	admin.DisplayUserData()

}

func getUserData(prompt string) string {
	fmt.Println(prompt)
	var data string
	fmt.Scanln(&data)
	return data
}
