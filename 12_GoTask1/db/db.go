package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not Open Database")
	}
	DB = db
	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)
	createTables()
}

func createTables() {
	// Create Projects table
	createProjectsTable := `
	CREATE TABLE IF NOT EXISTS projects (
	 id INTEGER PRIMARY KEY AUTOINCREMENT,
	 name TEXT NOT NULL,
	 description TEXT NOT NULL,
	 created_at DATETIME NOT NULL,
	 updated_at DATETIME NOT NULL
	)`
	_, err := DB.Exec(createProjectsTable)
	if err != nil {
		panic("Failed to created Projects Table")
	}

	//  Create Tasks table

	createTasksTable := `
	CREATE TABLE IF NOT EXISTS tasks(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL UNIQUE,
	description TEXT NOT NULL,
	status TEXT NOT NULL,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,
	project_id INTEGER,
	FOREIGN KEY(project_id) REFERENCES projects(id)
	)
	`
	_, err = DB.Exec(createTasksTable)
	if err != nil {
		panic("Failed to create tasks table")
		// panic(err)
	}

	// Create Team Members table
	createMembersTable := `
	CREATE TABLE IF NOT EXISTS members(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL
	)
	`
	_, err = DB.Exec(createMembersTable)
	if err != nil {
		panic("Failed to create tasks table")
		//  panic(err)
	}

	// Create Team Members table
	createTaskAssignmentTable := `
	CREATE TABLE IF NOT EXISTS task_assignments(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,
	task_id INTEGER,
	FOREIGN KEY(task_id) REFERENCES tasks(id),
	member_id INTEGER,
	FOREIGN KEY(member_id) REFERENCES members(id)
	)
	`
	_, err = DB.Exec(createTaskAssignmentTable)
	if err != nil {
		panic(err)
	}
}
