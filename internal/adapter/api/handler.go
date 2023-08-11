package api

import (
	"encoding/json"
	"github.com/tiagosberg/golang-example/pkg/http/result"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	. "github.com/tiagosberg/golang-example/internal/core/person"
)

func CreatePerson(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var cmd CreatePersonCommand
		if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
			result.InternalServerError(w)
			return
		}

		person, err := s.Create(cmd)
		if err != nil {
			result.BadRequest(w, err)
			return
		}

		result.Created(w, person)
	}
}

func FindPerson(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		person, _ := s.Find(FindPersonQuery{Id: id})
		if person == nil {
			result.NotFound(w)
			return
		}

		result.Ok(w, person)
	}
}

func GetPersons(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		persons, err := s.Get(GetPersonsQuery{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(persons) == 0 {
			result.NoContent(w)
			return
		}

		result.Ok(w, persons)
	}
}
