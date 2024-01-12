package validator

import (
	"fmt"
	"turing-machine/internal/domain"

	"github.com/samber/lo"
)

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

func SumHasParity(variants []Parity) []domain.Validator {
	return makeValidators[Parity](sumParityChecker{}, variants)
}
