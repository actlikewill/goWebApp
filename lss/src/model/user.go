package model

import (	
	"database/sql"	
	"fmt"
	"time"
)

const passwordSalt = "sdflasjdfladskfalf34lwk33290sadklasjdlkasdlvnasdklfladks"

type User struct {
	ID int 
	Email string
	Password string
	FirstName string
	LastName string
	LastLogin *time.Time
}

func Login(email, password string) (*User, error) {
	result := &User{}	
	row := db.QueryRow(`
		SELECT user_id, email 
		FROM users
		WHERE email = $1
		AND password = $2`, email, password)
	err := row.Scan(&result.ID, & result.Email)
	switch {
		case err == sql.ErrNoRows:
			return nil, fmt.Errorf("User not found")
		case err != nil:
			return nil, err
	}
	return result, nil
}