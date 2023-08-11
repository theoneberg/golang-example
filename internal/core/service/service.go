package service

import (
	"github.com/tiagosberg/golang-example/internal/core/domain"
	"github.com/tiagosberg/golang-example/internal/core/dto"
	"github.com/tiagosberg/golang-example/internal/core/port/driven"
)

type PersonService struct {
	repo driven.Repository
}

func (s PersonService) Create(cmd dto.CreatePersonCommand) (person *domain.Person, err error) {
	person = domain.NewPerson(cmd.Name, cmd.BirthDate)

	s.repo.Add(person)
	return
}

func (s PersonService) Find(query dto.FindPersonQuery) (person *domain.Person, err error) {
	person = s.repo.FindById(query.Id)
	return
}

func (s PersonService) Get(query dto.GetPersonsQuery) (persons []*domain.Person, err error) {
	persons = s.repo.Get()
	return
}

func NewPersonService(repo driven.Repository) PersonService {
	return PersonService{
		repo: repo,
	}
}
