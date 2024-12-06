package models

import (
	"fmt"
	"time"

	"github.com/gotask1/db"
)

type TaskAssignment struct {
	ID        int64     `json:"id"`
	TaskId    int64     `json:"task_id"`
	MemberId  int64     `json:"member_id"`
	CreatedAt time.Time `json:"creation-time"`
	UpdatedAt time.Time `json:"updation-time"`
}

func (ta TaskAssignment) Save() error {
	query := `
	INSERT INTO task_assignments(task_id, member_id, created_at, updated_at)
	VALUES(?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("Error==>>>>>>>", err)
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(ta.TaskId, ta.MemberId, ta.CreatedAt, ta.UpdatedAt)
	if err != nil {
		fmt.Println("Error==>>>>>>>", err)
		return err
	}

	id, err := result.LastInsertId()
	ta.ID = id

	fmt.Println("Error==>>>>>>>", err)
	return err
}

func (ta TaskAssignment) Update() error {
	query := `
	UPDATE task_assignments 
	SET task_id= ?, member_id= ?, created_at= ?, updated_at= ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(ta.TaskId, ta.MemberId, ta.CreatedAt, ta.UpdatedAt, ta.ID)
	if err != nil {
		return err
	}
	return nil
}

func (ta TaskAssignment) Delete() error {
	query := `
	Delete from task_assignments
	WHERE id = ?
	`
	_, err := db.DB.Exec(query, ta.ID)
	if err != nil {
		fmt.Println("error ===> ", err)
		return err
	}
	return nil
}

func GetAssignmentById(assignId int64) (*TaskAssignment, error) {
	query := `SELECT * FROM task_assignments WHERE id= ?`
	row := db.DB.QueryRow(query, assignId)

	var assignment TaskAssignment
	err := row.Scan(&assignment.ID, &assignment.CreatedAt, &assignment.UpdatedAt, &assignment.TaskId, &assignment.MemberId)
	if err != nil {
		fmt.Println("err=====>>>>>", err)
		return nil, err
	}
	return &assignment, nil
}

func GetAssignmentByTaskId(taskId int64) (*TaskAssignment, error) {
	query := `SELECT * FROM task_assignments WHERE task_id= ?`
	row := db.DB.QueryRow(query, taskId)

	var assignment TaskAssignment
	err := row.Scan(&assignment.ID, &assignment.CreatedAt, &assignment.UpdatedAt, &assignment.TaskId, &assignment.MemberId)
	if err != nil {
		fmt.Println("err=====>>>>>", err)
		return nil, err
	}
	return &assignment, nil
}

func GetAssignmentByMemberId(memberId int64) (*TaskAssignment, error) {
	query := `SELECT * FROM task_assignments WHERE member_id= ?`
	row := db.DB.QueryRow(query, memberId)

	var assignment TaskAssignment
	err := row.Scan(&assignment.ID, &assignment.CreatedAt, &assignment.UpdatedAt, &assignment.TaskId, &assignment.MemberId)
	if err != nil {
		fmt.Println("err=====>>>>>", err)
		return nil, err
	}
	return &assignment, nil
}

func GetAllAssignments() ([]TaskAssignment, error) {
	query := "SELECT * FROM task_assignments"
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println("err=====>>>>>", err)
		return nil, err

	}
	defer rows.Close()

	var assignents []TaskAssignment

	for rows.Next() {
		var assignent TaskAssignment
		err := rows.Scan(&assignent.ID, &assignent.CreatedAt, &assignent.UpdatedAt, &assignent.TaskId, &assignent.MemberId)
		if err != nil {
			fmt.Println("err=====>>>>>", err)
			return nil, err
		}
		assignents = append(assignents, assignent)
	}
	return assignents, nil
}
