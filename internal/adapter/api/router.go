package api

import (
	"github.com/go-chi/chi"
	"github.com/tiagosberg/golang-example/internal/core/person"
	"github.com/tiagosberg/golang-example/internal/storage/in-memory"
)

func MapRoutes(router *chi.Mux) {

	repo := inmemory.NewInMemoryRepository()

	svc := person.NewService(repo)

	router.Post("/persons", CreatePerson(svc))
	router.Get("/persons/{id}", FindPerson(svc))
	router.Get("/persons", GetPersons(svc))
}
