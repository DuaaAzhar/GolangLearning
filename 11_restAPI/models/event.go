package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding: "required"`
	Description string    `binding: "required"`
	Location    string    `binding: "required"`
	DateTime    time.Time `binding: "required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() {
	// will save the event in db
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
