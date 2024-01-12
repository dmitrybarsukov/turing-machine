package validator

import (
	"fmt"
	"turing-machine/internal/domain"

	"github.com/samber/lo"
)

type countOfNumberChecker struct {
	Number int

	count int
}

func (c countOfNumberChecker) Validate(code domain.Code) bool {
	return lo.Count(code[:], c.Number) == c.count
}

func (c countOfNumberChecker) String() string {
	return fmt.Sprintf("contains %d of '%d'", c.count, c.Number)
}

func (c countOfNumberChecker) WithValue(value int) domain.Validator {
	c.count = value
	return c
}

func CountOfNumber(value int, variants []int) []domain.Validator {
	return makeValidators[int](countOfNumberChecker{
		Number: value,
	}, variants)
}

type countOfParityChecker struct {
	Parity Parity

	count int
}

func (c countOfParityChecker) Validate(code domain.Code) bool {
	return lo.CountBy(code[:], func(i int) bool {
		return getParity(i) == c.Parity
	}) == c.count
}

func (c countOfParityChecker) String() string {
	return fmt.Sprintf("contains %d of %v numbers", c.count, c.Parity)
}

func (c countOfParityChecker) WithValue(value int) domain.Validator {
	c.count = value
	return c
}

func CountOfParity(parity Parity, variants []int) []domain.Validator {
	return makeValidators[int](countOfParityChecker{
		Parity: parity,
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
	return fmt.Sprintf("more %s numbers", c.result)
}

func (c parityCountComparator) WithValue(value Parity) domain.Validator {
	c.result = value
	return c
}

func HasMoreNumbersWithParity(variants []Parity) []domain.Validator {
	return makeValidators[Parity](parityCountComparator{}, variants)
}

type repetitionCounter struct {
	result int
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

func (c repetitionCounter) WithValue(value int) domain.Validator {
	c.result = value
	return c
}

func HasSomeRepeatingNumbers(variants []int) []domain.Validator {
	return makeValidators[int](repetitionCounter{}, variants)
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
