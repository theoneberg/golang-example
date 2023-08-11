package inmemory

import (
	"github.com/google/uuid"
	"github.com/tiagosberg/golang-example/internal/core/domain"
)

type InMemoryRepository struct {
	persons []*domain.Person
}

func (repo *InMemoryRepository) Add(person *domain.Person) {
	repo.persons = append(repo.persons, person)
}

func (repo *InMemoryRepository) FindById(id uuid.UUID) *domain.Person {
	for _, person := range repo.persons {
		if person.Id == id {
			return person
		}
	}

	return nil
}

func (repo *InMemoryRepository) FindByName(name string) *domain.Person {
	for _, person := range repo.persons {
		if person.Name == name {
			return person
		}
	}

	return nil
}

func (repo *InMemoryRepository) Get() []*domain.Person {
	return repo.persons
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		persons: make([]*domain.Person, 0),
	}
}
