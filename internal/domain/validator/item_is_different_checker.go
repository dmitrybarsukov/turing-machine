package validator

import (
	"fmt"
	"turing-machine/internal/domain"

	"github.com/samber/lo"
)

type itemIsDifferentChecker struct {
	Result Compare

	item domain.CodeItem
}

func (c itemIsDifferentChecker) Validate(code domain.Code) bool {
	itemsExceptTarget := lo.Filter(code[:], func(_ int, idx int) bool {
		return idx != c.item.Index()
	})

	isOk := true
	for _, item := range itemsExceptTarget {
		if compare(code.Get(c.item), item) != c.Result {
			isOk = false
		}
	}

	return isOk
}

func (c itemIsDifferentChecker) String() string {
	return fmt.Sprintf("%v %v others", c.item, c.Result)
}

func (c itemIsDifferentChecker) WithValue(value domain.CodeItem) domain.Validator {
	c.item = value
	return c
}

func OneItemIsDifferent(compareResult Compare, variants []domain.CodeItem) []domain.Validator {
	return makeValidators[domain.CodeItem](itemIsDifferentChecker{Result: compareResult}, variants)
}
