package api

import (
	"encoding/json"
	"net/http"

	"github.com/tiagosberg/golang-example/pkg/http/result"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/tiagosberg/golang-example/internal/core/dto"
	"github.com/tiagosberg/golang-example/internal/core/port/driving"
)

func CreatePerson(s driving.PersonService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var cmd dto.CreatePersonCommand
		if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
			result.BadRequest(w, err)
			return
		}

		person, err := s.Create(cmd)
		if err != nil {
			result.InternalServerError(w)
			return
		}

		result.Created(w, person)
	}
}

func FindPerson(s driving.PersonService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			result.BadRequest(w, err)
			return
		}

		person, err := s.Find(dto.FindPersonQuery{Id: id})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if person == nil {
			result.NotFound(w)
			return
		}

		result.Ok(w, person)
	}
}

func GetPersons(s driving.PersonService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		persons, err := s.Get(dto.GetPersonsQuery{})
		if err != nil {
			result.BadRequest(w, err)
			return
		}

		if len(persons) == 0 {
			result.NoContent(w)
			return
		}

		result.Ok(w, persons)
	}
}
