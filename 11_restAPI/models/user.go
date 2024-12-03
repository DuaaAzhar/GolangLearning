package models

import "github.com/restAPI/db"

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

	result, err := stmt.Exec(u.Email, u.Password)
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

func GetAllUsers() {

}

func GetUserByID() {

}
