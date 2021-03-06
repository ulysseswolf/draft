package model

import (
	"fmt"
	"time"

	"app/shared/database"
)

// *****************************************************************************
// User
// *****************************************************************************

// User table contains the information for each user
type User struct {
	ID        uint32    `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	StatusID  uint8     `db:"status_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Deleted   uint8     `db:"deleted"`
}

// UserStatus table contains every possible user status (active/inactive)
type UserStatus struct {
	ID        uint8     `db:"id"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Deleted   uint8     `db:"deleted"`
}

// UserID returns the user id
func (u *User) UserID() string {
	r := ""

	r = fmt.Sprintf("%v", u.ID)
	return r
}

// UserByEmail gets user information from email
func UserByEmail(email string) (User, error) {
	var err error

	result := User{}

	err = database.SQL.Get(&result, "SELECT id, password, status_id, first_name FROM user WHERE email = ? LIMIT 1", email)
	return result, standardizeError(err)
}

// UserCreate creates user
func UserCreate(firstName, lastName, email, password string) error {
	var err error

	_, err = database.SQL.Exec("INSERT INTO user (first_name, last_name, email, password) VALUES (?,?,?,?)", firstName, lastName, email, password)

	return standardizeError(err)
}
