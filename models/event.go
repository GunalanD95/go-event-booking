package models

import (
	"event_booking/db"
	"fmt"
	"time"
)

type Event struct {
	Id       int
	Name     string    `binding:"required"`
	Location string    `binding:"required"`
	Date     time.Time `binding:"required"`
	Price    int       `binding:"required"`
	UserId   int
}

type EventRegistration struct {
	Name    string `binding:"required"`
	UserId  int    `binding:"required"`
	EventId int    `binding:"required"`
}

func (e *EventRegistration) Save() error {
	query := `
	INSERT INTO events (name, user_id, event_id) 
	VALUES (?, ?, ?);
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %v", err)
	}
	defer statement.Close()
	result, err := statement.Exec(e.Name, e.UserId, e.EventId)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}
	_, err = result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to fetch last inserted ID: %v", err)
	}
	return nil
}

func (e Event) Save() error {
	query := `
	INSERT INTO EVENTS (name, location, date, price, user_id) VALUES (?, ?, ?, ?, ?);
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(e.Name, e.Location, e.Date, e.Price, e.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.Id = int(id)
	fmt.Printf("Event saved to database: %v\n", e)
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `
	SELECT * FROM EVENTS;
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Location, &event.Date, &event.Price, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(event_id int) (Event, error) {
	query := `
	SELECT *
	FROM EVENTS as e
	WHERE e.id = ?
	`
	var event Event
	row := db.DB.QueryRow(query, event_id)
	err := row.Scan(&event.Id, &event.Name, &event.Location, &event.Date, &event.Price, &event.UserId)
	if err != nil {
		// Return other types of errors
		return event, fmt.Errorf("error querying event: %w", err)
	}

	return event, nil

}

func UpdateEventById(event_id int, updatedEvent Event) error {
	query := `
	UPDATE EVENTS
	SET name = ?, location = ?, date = ?, price = ?, user_id = ?
	WHERE id = ?
	`
	// Execute the update query
	_, err := db.DB.Exec(query, updatedEvent.Name, updatedEvent.Location, updatedEvent.Date, updatedEvent.Price, updatedEvent.UserId, event_id)
	if err != nil {
		return fmt.Errorf("error updating event: %w", err)
	}

	return nil
}

func DeleteEventById(event_id int) error {
	query := `
	DELETE FROM EVENTS
	WHERE id = ?
	`
	// Execute the update query
	_, err := db.DB.Exec(query, event_id)
	if err != nil {
		return fmt.Errorf("error updating event: %w", err)
	}

	return nil
}
