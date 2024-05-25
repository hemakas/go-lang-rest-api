package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

// empty slice
var events = []Event{}

// save event
func (e Event) Save() {
	events = append(events, e)
}

// get all events
func GetAllEvents() []Event {
	return events
}
