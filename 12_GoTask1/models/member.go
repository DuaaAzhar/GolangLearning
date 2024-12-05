package models

import (
	"fmt"
	"time"

	"github.com/gotask1/db"
)

type Member struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"creation-time"`
	UpdatedAt time.Time `json:"updation-time"`
}

func (m Member) Save() error {
	query := `
	INSERT INTO members(name, email, created_at, updated_at)
	VALUES(?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(m.Name, m.Email, m.CreatedAt, m.UpdatedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	m.ID = id

	return err
}

func (m Member) Update() error {
	query := `
	UPDATE members 
	SET name= ?, email= ?, created_at= ?, updated_at= ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(m.Name, m.Email, m.CreatedAt, m.UpdatedAt, m.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m Member) Delete() error {
	query := `
	Delete from members
	WHERE id = ?
	`
	_, err := db.DB.Exec(query, m.ID)
	if err != nil {
		fmt.Println("error ===> ", err)
		return err
	}
	return nil
}

func GetMemberById(memberId int64) (*Member, error) {
	query := `SELECT * FROM members WHERE id= ?`
	row := db.DB.QueryRow(query, memberId)

	var member Member
	err := row.Scan(&member.ID, &member.Name, &member.Email, &member.CreatedAt, &member.UpdatedAt)
	if err != nil {
		fmt.Println("err=====>>>>>", err)
		return nil, err
	}
	return &member, nil
}

func GetAllMembers() ([]Member, error) {
	query := "SELECT * FROM members"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []Member

	for rows.Next() {
		var member Member
		err := rows.Scan(&member.ID, &member.Name, &member.Email, &member.CreatedAt, &member.UpdatedAt)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
}
