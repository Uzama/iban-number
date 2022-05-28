package http

import (
	"net/http"

	"github.com/Uzama/iban-number/validator"
)

type handler struct {
	validator validator.Validator
}

func NewHandler() handler {

	return handler{
		validator: validator.NewValidator(),
	}
}

func (h handler) ValidateIBANNumber(w http.ResponseWriter, r *http.Request) {

}
