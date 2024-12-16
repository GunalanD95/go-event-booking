package models

import (
	"event_booking/db"
	"fmt"
)

type User struct {
	Id       int
	Name     string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO USERS (name, email, password) VALUES (?, ?, ?);
	`
	statement, err := db.DB.Prepare(query)
	defer statement.Close()
	if err != nil {
		return err
	}
	_, err = statement.Exec(u.Name, u.Email, u.Password)
	if err != nil {
		return err
	}
	fmt.Printf("User saved to database: %v\n", u)
	return nil
}
