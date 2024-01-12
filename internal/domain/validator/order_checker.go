package validator

import (
	"fmt"
	"turing-machine/internal/domain"
)

type orderChecker struct {
	result Order
}

func (c orderChecker) Validate(code domain.Code) bool {
	return getOrder(code[:]) == c.result
}

func (c orderChecker) String() string {
	return fmt.Sprintf("code is %v", c.result)
}

func (c orderChecker) WithValue(value Order) domain.Validator {
	c.result = value
	return c
}

func CodeIsOrdered(variants []Order) []domain.Validator {
	return makeValidators[Order](orderChecker{}, variants)
}
