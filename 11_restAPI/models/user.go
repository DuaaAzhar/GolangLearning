package models

import (
	"errors"

	"github.com/restAPI/db"
	"github.com/restAPI/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO users( email, password)
	VALUES(?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err

}

func (u User) Update() error {
	query := `
	UPDATE users 
	SET email= ?, password= ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Email, u.Password, u.ID)
	return err
}

func (u User) Delete() error {
	query := `
	Delete users
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.ID)
	return err
}

func GetAllUsers() ([]User, error) {
	query := `SELECT * FROM USERS`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserByID(userId int64) (*User, error) {
	query := `
	SELECT * FROM users 
	WHERE id= ?`

	row := db.DB.QueryRow(query, userId)
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) ValidateUser() error {
	query := `SELECT id, password FROM users WHERE email= ?`
	row := db.DB.QueryRow(query, u.Email)

	var savedPassword string
	err := row.Scan(&u.ID, &savedPassword)

	if err != nil {
		return errors.New("invalid credentials")
	}
	valid := utils.CheckPasswordHash(u.Password, savedPassword)
	if !valid {
		return errors.New("invalid credentials")
	}

	return nil

}
