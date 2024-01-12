package validator

import (
	"fmt"
	"turing-machine/internal/domain"

	"github.com/samber/lo"
)

type majorParityChecker struct {
	Parity Parity
}

func (c majorParityChecker) Validate(code domain.Code) bool {
	expectedCount := lo.CountBy(code[:], func(i int) bool {
		return getParity(i) == c.Parity
	})
	unexpectedCount := len(code) - expectedCount
	return expectedCount > unexpectedCount
}

func (c majorParityChecker) String() string {
	return fmt.Sprintf("%s majority", c.Parity)
}

type hasSameNumbersPairChecker struct {
	Result bool
}

func (c hasSameNumbersPairChecker) Validate(code domain.Code) bool {
	grouped := lo.GroupBy(code[:], func(it int) int {
		return it
	})

	hasPair := lo.ContainsBy(lo.Values(grouped), func(ints []int) bool {
		return len(ints) == 2
	})

	return hasPair == c.Result
}

func (c hasSameNumbersPairChecker) String() string {
	if c.Result {
		return "has pair"
	}

	return "no pair"
}

type itemIsOutstandingChecker struct {
	Item    domain.CodeItem
	Compare Compare
}

func (c itemIsOutstandingChecker) Validate(code domain.Code) bool {
	itemsExceptTarget := lo.Filter(code[:], func(_ int, idx int) bool {
		return idx != c.Item.Index()
	})

	isOk := true
	for _, item := range itemsExceptTarget {
		if compare(code.Get(c.Item), item) != c.Compare {
			isOk = false
		}
	}

	return isOk
}

func (c itemIsOutstandingChecker) String() string {
	return fmt.Sprintf("%v %v others", c.Item, c.Compare)
}

type orderStrictChecker struct {
	Order Order
}

func (c orderStrictChecker) Validate(code domain.Code) bool {
	return getOrderStrict(code[:]) == c.Order
}

func (c orderStrictChecker) String() string {
	return fmt.Sprintf("code is %v", c.Order)
}
