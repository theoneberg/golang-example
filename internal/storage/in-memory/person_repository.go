package inmemory

import (
	"github.com/google/uuid"

	personpkg "github.com/tiagosberg/golang-example/internal/core/person"
)

type inMemoryRepository struct {
	persons []*personpkg.Person
}

func (repo *inMemoryRepository) Add(person *personpkg.Person) {
	repo.persons = append(repo.persons, person)
}

func (repo *inMemoryRepository) FindById(id uuid.UUID) *personpkg.Person {

	for _, person := range repo.persons {
		if person.Id == id {
			return person
		}
	}

	return nil
}

func (repo *inMemoryRepository) FindByName(name string) *personpkg.Person {

	for _, person := range repo.persons {
		if person.Name == name {
			return person
		}
	}

	return nil
}

func (repo *inMemoryRepository) Get() []*personpkg.Person {
	return repo.persons
}

func NewInMemoryRepository() personpkg.Repository {
	return &inMemoryRepository{
		persons: make([]*personpkg.Person, 0),
	}
}
