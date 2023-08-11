package driving

import (
	"github.com/tiagosberg/golang-example/internal/core/domain"
	"github.com/tiagosberg/golang-example/internal/core/dto"
)

type PersonService interface {
	Create(dto.CreatePersonCommand) (*domain.Person, error)
	Find(dto.FindPersonQuery) (*domain.Person, error)
	Get(dto.GetPersonsQuery) ([]*domain.Person, error)
}
