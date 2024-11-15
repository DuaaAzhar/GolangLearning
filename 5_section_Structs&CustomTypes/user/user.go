// struct User
// struct Admin embedding struct User

package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	dob       string
	createdAt time.Time
}

// Struct Emebedding
type Admin struct {
	email    string
	password string
	User     //if we wanna directly embed the struct like this
	// OR
	// U User
}

// Defining Method having receiver param of type user struct
func (u User) DisplayUserData() {
	fmt.Println(u.firstName, u.lastName, u.dob)
}

// Creating a Receiver param method that mutates the struct value
// if we pass the value like this (u user) it will not mutate struct bcz receiver methods pass only copies not original value
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(userFirstName, userLastName, userBirthDate string) (*User, error) {
	if userFirstName == "" || userLastName == "" {
		return nil, errors.New("First Name or Last Name cant be empty")
	}
	return &User{
		firstName: userFirstName,
		lastName:  userLastName,
		dob:       userBirthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password, firstName, lastName, dob string) Admin {
	return Admin{ // in case of defining direct embedded struct
		email:    email,
		password: password,
		User: User{
			firstName: firstName,
			lastName:  lastName,
			dob:       dob,
			createdAt: time.Now(),
		},
	}
	// return Admin{
	// 	email:    email,
	// 	password: password,
	// 	U: User{
	// 		firstName: firstName,
	// 		lastName:  lastName,
	// 		dob:       dob,
	// 		createdAt: time.Now(),
	// 	},
	// }
}
