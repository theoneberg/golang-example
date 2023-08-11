package api

import (
	"github.com/go-chi/chi"
	"github.com/tiagosberg/golang-example/internal/core/port/driving"
)

func MapRoutes(router *chi.Mux, service driving.PersonService) {
	router.Post("/persons", CreatePerson(service))
	router.Get("/persons/{id}", FindPerson(service))
	router.Get("/persons", GetPersons(service))
}
