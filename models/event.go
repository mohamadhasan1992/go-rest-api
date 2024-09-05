package models

import (
	"fmt"
	"time"

	"github.com/mohamadhasan1992/go-rest-api.git/db"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	fmt.Println("event body", e)
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("****")
		return err
	}
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	defer stmt.Close()
	if err != nil {
		fmt.Println("////****///")
		return err
	}
	Id, err := result.LastInsertId()
	e.Id = Id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventDetail(eventId int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, eventId)
	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name =?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.Id)
	return err
}

func (event Event) Delete() error {
	query := `
	DELETE FROM events
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.Id)
	return err
}

func (event *Event) Register(userId int64) error {
	query := `
	INSERT INTO registerations(user_id, event_id)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	result, err := stmt.Exec(userId, event.Id)
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	Id, err := result.LastInsertId()
	event.Id = Id
	return err
}

func (event Event) DeleteRegister(userId int64) error {
	query := "DELETE FROM registerations WHERE user_id = ? AND event_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = stmt.Exec(userId, event.Id)
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}
