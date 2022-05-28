package http

import (
	"log"
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

	// parse param values
	ibanNumber := r.FormValue("iban_number")

	// validate the value whether it is empty or not
	if ibanNumber == "" {
		err := "iban number not given"

		log.Println(err)

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err))
		return
	}

	log.Printf("request:[iban-number:%s]", ibanNumber)

	// check whether the iban number is valid or not
	isValid, err := h.validator.Validate(ibanNumber)
	if err != nil {
		log.Println(err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}

	log.Printf("response:[is-valid:%v]", isValid)

	// encode and send the response
	response := Respose{IsValid: isValid}

	w.WriteHeader(http.StatusOK)
	w.Write(response.Encode())
}
