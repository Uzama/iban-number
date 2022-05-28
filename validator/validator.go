package validator

type validator struct {
}

func NewValidator() Validator {
	return &validator{}
}

func (v validator) Validate(iban string) bool {
	return false
}
