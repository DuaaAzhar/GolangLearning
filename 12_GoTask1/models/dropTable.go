package models

import (
	"fmt"

	"github.com/gotask1/db"
)

func DeleteTables() error {
	query := `
	DROP TABLE IF EXISTS tasks;
	DROP TABLE IF EXISTS members;
	DROP TABLE IF EXISTS projects;
	DROP TABLE IF EXISTS task_assignments
	`

	_, err := db.DB.Exec(query)

	if err != nil {
		fmt.Println("err=====>>>>>", err)
		return err
	}
	return nil
}
