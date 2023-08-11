package person

import (
	"github.com/google/uuid"
	"github.com/tiagosberg/golang-example/pkg/date"
)

type CreatePersonCommand struct {
	Name      string    `json:"name"`
	BirthDate date.Date `json:"birth_date"`
}

type FindPersonQuery struct {
	Id uuid.UUID
}

type GetPersonsQuery struct{}
