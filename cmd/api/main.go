package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/tiagosberg/golang-example/internal/adapter/api"
)

func main() {
	router := chi.NewMux()

	api.MapRoutes(router)

	log.Fatalln(http.ListenAndServe(":8080", router))
}
