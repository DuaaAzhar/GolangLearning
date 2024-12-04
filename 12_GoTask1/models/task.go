package models

import (
	"fmt"
	"time"

	"github.com/gotask1/db"
)

type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Project_ID  int64     `json:"project_id"`
	CreatedAt   time.Time `json:"creation-time"`
	UpdatedAt   time.Time `json:"updation-time"`
}

func (t Task) Save() error {
	query := `
	INSERT INTO tasks(title, description, status, project_id, created_at, updated_at)
	VALUES(?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(t.Title, t.Description, t.Status, t.Project_ID, t.CreatedAt, t.UpdatedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	t.ID = id

	return err
}

func (t Task) Update() error {
	query := `
	UPDATE tasks 
	SET title= ?, description= ?, status= ? , project_id= ?, created_at= ?, updated_at= ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.Title, t.Description, t.Status, t.Project_ID, t.CreatedAt, t.UpdatedAt, t.ID)
	if err != nil {
		return err
	}
	return nil
}

func (t Task) Delete() error {
	query := `
	Delete from tasks
	WHERE id = ?
	`
	_, err := db.DB.Exec(query, t.ID)
	if err != nil {
		fmt.Println("error ===> ", err)
		return err
	}
	return nil
}

func GetTaskById(projectId int64) (*Task, error) {
	query := `SELECT * FROM tasks WHERE id= ?`
	row := db.DB.QueryRow(query, projectId)

	var task Task
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		fmt.Println("err=====>>>>>", err)
		return nil, err
	}
	return &task, nil
}

func GetAllTasks() ([]Task, error) {
	query := "SELECT * FROM tasks"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
