package validator

import (
	"fmt"
	"turing-machine/internal/domain"

	"github.com/samber/lo"
)

type countOfPredicateChecker struct {
	Predicate   func(i int) bool
	Description string

	count Count
}

func (c countOfPredicateChecker) Validate(code domain.Code) bool {
	return lo.CountBy(code[:], c.Predicate) == int(c.count)
}

func (c countOfPredicateChecker) String() string {
	return fmt.Sprintf("contains %d of %s", c.count, c.Description)
}

func (c countOfPredicateChecker) WithValue(value Count) domain.Validator {
	c.count = value
	return c
}

func CountOfNumber(value int, variants []Count) []domain.Validator {
	return makeValidators[Count](countOfPredicateChecker{
		Predicate: func(i int) bool {
			return i == value
		},
		Description: fmt.Sprintf("'%d'", value),
	}, variants)
}

func CountOfParity(parity Parity, variants []Count) []domain.Validator {
	return makeValidators[Count](countOfPredicateChecker{
		Predicate: func(i int) bool {
			return getParity(i) == parity
		},
		Description: fmt.Sprintf("%v numbers", parity),
	}, variants)
}

type parityCountComparator struct {
	result Parity
}

func (c parityCountComparator) Validate(code domain.Code) bool {
	expectedCount := lo.CountBy(code[:], func(i int) bool {
		return getParity(i) == c.result
	})
	unexpectedCount := len(code) - expectedCount
	return expectedCount > unexpectedCount
}

func (c parityCountComparator) String() string {
	return fmt.Sprintf("has more %s than %s numbers", c.result, c.result.not())
}

func (c parityCountComparator) WithValue(value Parity) domain.Validator {
	c.result = value
	return c
}

func HasMoreNumbersWithParity(variants []Parity) []domain.Validator {
	return makeValidators[Parity](parityCountComparator{}, variants)
}

type repetitionCounter struct {
	result Count
}

func (c repetitionCounter) Validate(code domain.Code) bool {
	grouped := lo.GroupBy(code[:], func(it int) int {
		return it
	})

	maxCount := len(lo.MaxBy(lo.Values(grouped), func(ints []int, ints2 []int) bool {
		return len(ints) > len(ints2)
	}))
	repCount := maxCount - 1

	return repCount == int(c.result)
}

func (c repetitionCounter) String() string {
	if c.result == 0 {
		return "has no repetitions"
	}

	return fmt.Sprintf("has %d repeating numbers", c.result+1)
}

func (c repetitionCounter) WithValue(value Count) domain.Validator {
	c.result = value
	return c
}

func HasSomeRepeatingNumbers(variants []Count) []domain.Validator {
	return makeValidators[Count](repetitionCounter{}, variants)
}

type hasPairOfSameNumbersChecker struct {
	result bool
}

func (c hasPairOfSameNumbersChecker) Validate(code domain.Code) bool {
	grouped := lo.GroupBy(code[:], func(it int) int {
		return it
	})

	hasPair := lo.ContainsBy(lo.Values(grouped), func(ints []int) bool {
		return len(ints) == 2
	})

	return hasPair == c.result
}

func (c hasPairOfSameNumbersChecker) String() string {
	if c.result {
		return "has pair of same numbers"
	}

	return "does not have pair of same numbers"
}

func (c hasPairOfSameNumbersChecker) WithValue(value bool) domain.Validator {
	c.result = value
	return c
}

func PairOfNumbersExist(variants []bool) []domain.Validator {
	return makeValidators[bool](hasPairOfSameNumbersChecker{}, variants)
}
