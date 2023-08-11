package person

import (
	"github.com/google/uuid"
	"github.com/tiagosberg/golang-example/pkg/date"
)

type Person struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	BirthDate date.Date `json:"birth_date"`
}

func NewPerson(name string, birthDate date.Date) *Person {
	return &Person{
		Id:        uuid.New(),
		Name:      name,
		BirthDate: birthDate,
	}
}
