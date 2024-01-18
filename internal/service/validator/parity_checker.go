package validator

import (
	"fmt"

	"github.com/dmitrybarsukov/turing-machine/internal/domain"

	"github.com/samber/lo"
)

type itemParityChecker struct {
	Item   domain.CodeItem
	Parity Parity
}

func (c itemParityChecker) Validate(code domain.Code) bool {
	return getParity(code.Get(c.Item)) == c.Parity
}

func (c itemParityChecker) String() string {
	return fmt.Sprintf("%v is %v", c.Item, c.Parity)
}

type sumParityChecker struct {
	Parity Parity
}

func (c sumParityChecker) Validate(code domain.Code) bool {
	return getParity(lo.SumBy[int, int](code[:], func(i int) int {
		return i
	})) == c.Parity
}

func (c sumParityChecker) String() string {
	return fmt.Sprintf("sum is %v", c.Parity)
}
