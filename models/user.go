package models

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

type User struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    []byte    `json:"password"`
	Group       Group     `json:"group"`
	IsActive    bool      `json:"is_active"`
	PhoneNumber string    `json:"phone_number"`
	Photo       string    `json:"photo"`
	Biography   string    `json:"biography"`
	Birthdate   time.Time `json:"birthdate,string,omitempty"`
	Gender      string    `json:"gender"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserNull struct {
	PhoneNumber sql.NullString
	Photo       sql.NullString
	Biography   sql.NullString
	Birthdate   mysql.NullTime
}

func UserValidation(user User) (User, error) {
	// todo : sanitation and validation of user model
	return user, nil
}
