package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	// create a router
	r := mux.NewRouter()

	// initialize handler
	handler := NewHandler()

	// define APIs
	r.HandleFunc("/validate", handler.ValidateIBANNumber).Methods(http.MethodGet)

	return r
}
