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

	ibanNumber := r.FormValue("iban_number")

	if ibanNumber == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("iban number not given"))
		return
	}

	isValid, err := h.validator.Validate(ibanNumber)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}

	response := Respose{IsValid: isValid}

	w.WriteHeader(http.StatusOK)
	w.Write(response.Encode())
}
