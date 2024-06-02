package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	BirthDate time.Time `json:"birth_date"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
}

func NewUser(name string, birthDate time.Time, email string, password string) (*User, error) {
	usr := &User{
		ID:        uuid.New(),
		BirthDate: birthDate,
		Email:     email,
		Password:  password,
		Name:      name,
	}
	return usr, nil
}


