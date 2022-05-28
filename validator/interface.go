package validator

type Validator interface {
	Validate(iban string) (bool, error)
}
