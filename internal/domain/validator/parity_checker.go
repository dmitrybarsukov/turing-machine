package validator

import (
	"fmt"
	"turing-machine/internal/domain"
)

type itemParityChecker struct {
	Item domain.CodeItem

	result Parity
}

func (c itemParityChecker) Validate(code domain.Code) bool {
	return getParity(code.Get(c.Item)) == c.result
}

func (c itemParityChecker) String() string {
	return fmt.Sprintf("%v is %v", c.Item, c.result)
}

func (c itemParityChecker) WithValue(value Parity) domain.Validator {
	c.result = value
	return c
}

func ItemHasParity(item domain.CodeItem, variants []Parity) []domain.Validator {
	return makeValidators[Parity](itemParityChecker{Item: item}, variants)
}
