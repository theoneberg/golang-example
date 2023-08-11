package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/tiagosberg/golang-example/internal/adapter/inbound/api"
	"github.com/tiagosberg/golang-example/internal/adapter/outbound/storage/inmemory"
	"github.com/tiagosberg/golang-example/internal/core/service"
)

func main() {
	router := chi.NewMux()

	s := service.NewPersonService(inmemory.NewInMemoryRepository())

	api.MapRoutes(router, s)

	log.Fatalln(http.ListenAndServe(":8080", router))
}
