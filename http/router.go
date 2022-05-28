package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	handler := NewHandler()

	r.HandleFunc("/validate", handler.ValidateIBANNumber).Methods(http.MethodGet)

	return r
}
