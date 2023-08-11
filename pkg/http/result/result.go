package result

import (
	"encoding/json"
	"log"
	"net/http"
)

func status(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}

func respondJson(w http.ResponseWriter, statusCode int, object any) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	objectJson, err := json.Marshal(object)
	if err != nil {
		log.Fatalln("There was an error while marshal struct")
	}

	w.Write(objectJson)
}

func Ok(w http.ResponseWriter, object any) {
	respondJson(w, http.StatusOK, object)
}

func Created(w http.ResponseWriter, object any) {
	respondJson(w, http.StatusCreated, object)
}

func NoContent(w http.ResponseWriter) {
	status(w, http.StatusNoContent)
}

func BadRequest(w http.ResponseWriter, object any) {
	respondJson(w, http.StatusBadRequest, object)
}

func NotFound(w http.ResponseWriter) {
	status(w, http.StatusNotFound)
}

func InternalServerError(w http.ResponseWriter) {
	status(w, http.StatusInternalServerError)
}
