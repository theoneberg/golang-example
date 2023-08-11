package person

import "github.com/google/uuid"

type Repository interface {
	Add(*Person)
	FindById(uuid.UUID) *Person
	FindByName(string) *Person
	Get() []*Person
}
