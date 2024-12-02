package models

import (
	"time"

	"github.com/restAPI/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"date-time"`
	UserID      int       `json:"user-id"`
}

var events = []Event{}

func (e Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES(?,?,?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err

}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * from events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)

	}
	return events, nil
}
