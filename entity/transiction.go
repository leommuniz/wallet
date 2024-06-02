package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transiction struct {
	ID          uuid.UUID `json:"id"`
	Value    float64 `json:"value"`
	CreateIn    time.Time `json:"create_in"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Recurrent   bool      `json:"recurrent"`
}

func NewTransiction(value float64,recurrent bool, transictionType string, description string) *Transiction {
	Transiction := &Transiction{
		ID:          uuid.New(),
		Value:    value,
		CreateIn:    time.Now(),
		Type:        transictionType,
		Description: description,
		Recurrent:   recurrent,
	}
	return Transiction
}
