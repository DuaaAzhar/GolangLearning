package models

import (
	"fmt"
	"time"

	"github.com/gotask1/db"
)

type Project struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"creation-time"`
	UpdatedAt   time.Time `json:"updation-time"`
}

func (p Project) Save() error {
	query := `
	INSERT INTO projects(name, description, created_at, updated_at)
	VALUES(?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(p.Name, p.Description, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	p.ID = id

	return err
}

func (p Project) Update() error {
	query := `
	UPDATE projects 
	SET name= ?, description= ?, created_at= ?, updated_at= ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Name, p.Description, p.CreatedAt, p.UpdatedAt, p.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p Project) Delete() error {
	query := `
	Delete from projects
	WHERE id = ?
	`
	_, err := db.DB.Exec(query, p.ID)
	if err != nil {
		fmt.Println("error ===> ", err)
		return err
	}
	return nil
}

func GetProjectById(projectId int64) (*Project, error) {
	query := `SELECT * FROM projects WHERE id= ?`
	row := db.DB.QueryRow(query, projectId)

	var project Project
	err := row.Scan(&project.ID, &project.Name, &project.Description, &project.CreatedAt, &project.UpdatedAt)
	if err != nil {
		fmt.Println("err=====>>>>>", err)
		return nil, err
	}
	return &project, nil
}

func GetAllProjects() ([]Project, error) {
	query := "SELECT * FROM projects"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []Project

	for rows.Next() {
		var project Project
		err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.CreatedAt, &project.UpdatedAt)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}
