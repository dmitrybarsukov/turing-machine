package domain

type Validator interface {
	Validate(code Code) bool
}
