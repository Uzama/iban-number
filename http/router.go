package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("validate", nil).Methods(http.MethodGet)

	return r
}
