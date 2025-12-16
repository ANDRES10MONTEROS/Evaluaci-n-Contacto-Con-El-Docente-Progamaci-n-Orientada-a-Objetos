package models

import (
	"errors"
	"time"
)

type User struct {
	id         string
	name       string
	email      string
	subscribed bool
	createdAt  time.Time
}

type UserResponse struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Subscribed bool      `json:"subscribed"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewUser(id, name, email string) (*User, error) {
	if id == "" || name == "" || email == "" {
		return nil, errors.New("id, name and email are required")
	}
	return &User{
		id:         id,
		name:       name,
		email:      email,
		subscribed: false,
		createdAt:  time.Now(),
	}, nil
}

func (u *User) GetID() string           { return u.id }
func (u *User) GetName() string         { return u.name }
func (u *User) GetEmail() string        { return u.email }
func (u *User) GetSubscribed() bool     { return u.subscribed }
func (u *User) GetCreatedAt() time.Time { return u.createdAt }

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:         u.id,
		Name:       u.name,
		Email:      u.email,
		Subscribed: u.subscribed,
		CreatedAt:  u.createdAt,
	}
}

func (u *User) SetName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	u.name = name
	return nil
}

func (u *User) SetEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}
	u.email = email
	return nil
}
