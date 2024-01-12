package validator

import (
	"fmt"
	"turing-machine/internal/domain"

	"github.com/samber/lo"
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

func ItemHasParity(item domain.CodeItem) []domain.Validator {
	return makeValidators[Parity](itemParityChecker{Item: item}, parityVariants)
}

type sumParityChecker struct {
	result Parity
}

func (c sumParityChecker) Validate(code domain.Code) bool {
	return getParity(lo.SumBy[int, int](code[:], func(i int) int {
		return i
	})) == c.result
}

func (c sumParityChecker) String() string {
	return fmt.Sprintf("sum of numbers is %v", c.result)
}

func (c sumParityChecker) WithValue(value Parity) domain.Validator {
	c.result = value
	return c
}

func SumHasParity() []domain.Validator {
	return makeValidators[Parity](sumParityChecker{}, parityVariants)
}
