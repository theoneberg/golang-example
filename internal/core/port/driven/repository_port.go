package driven

import (
	"github.com/google/uuid"
	"github.com/tiagosberg/golang-example/internal/core/domain"
)

type Repository interface {
	Add(*domain.Person)
	FindById(uuid.UUID) *domain.Person
	FindByName(string) *domain.Person
	Get() []*domain.Person
}
