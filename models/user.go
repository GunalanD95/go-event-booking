package models

import (
	"database/sql"
	"event_booking/db"
	"fmt"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"` // Hashed password
}

func (u *User) Save() error {
	query := `
	INSERT INTO USERS (name, email, password) VALUES (?, ?, ?);
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(u.Name, u.Email, u.Password)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.Id = int(id)
	fmt.Printf("User saved to database: %v\n", u)
	return nil
}

func GetUserByEmail(email string) (*User, error) {
	query := `
	SELECT *
	FROM USERS as u
	WHERE u.email = ?
	`
	var user User
	err := db.DB.QueryRow(query, email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with email %s", email)
		}
		return nil, fmt.Errorf("error querying user: %w", err)
	}
	return &user, nil
}
