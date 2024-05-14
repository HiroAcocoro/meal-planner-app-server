package types

import (
	"time"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id string) (*User, error)
	CreateUser(User) error
}

type User struct {
	Id       string    `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	Created  time.Time `json:"created"`
}

type SignupUserPayload struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=16"`
}

type LoginUserPayload struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type contextKey string
const UserKey contextKey = "userId"