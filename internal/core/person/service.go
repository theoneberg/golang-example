package person

type Service interface {
	Create(CreatePersonCommand) (*Person, error)
	Find(FindPersonQuery) (*Person, error)
	Get(GetPersonsQuery) ([]*Person, error)
}

type service struct {
	repo Repository
}

func (s service) Create(cmd CreatePersonCommand) (*Person, error) {

	person := NewPerson(cmd.Name, cmd.BirthDate)

	s.repo.Add(person)

	return person, nil
}

func (s service) Find(query FindPersonQuery) (*Person, error) {

	person := s.repo.FindById(query.Id)

	return person, nil
}

func (s service) Get(query GetPersonsQuery) ([]*Person, error) {

	persons := s.repo.Get()

	return persons, nil
}

func NewService(repo Repository) Service {
	return service{
		repo: repo,
	}
}
