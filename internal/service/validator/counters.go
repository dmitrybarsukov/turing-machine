package validator

import (
	"fmt"
	"strconv"
	"turing-machine/internal/domain"

	"github.com/samber/lo"
)

type numberChecker interface {
	Check(value int) bool
}

type equalityChecker struct {
	Number int
}

func (c equalityChecker) Check(value int) bool {
	return value == c.Number
}

func (c equalityChecker) String() string {
	return strconv.Itoa(c.Number)
}

type parityChecker struct {
	Parity Parity
}

func (c parityChecker) Check(value int) bool {
	return getParity(value) == c.Parity
}

func (c parityChecker) String() string {
	return c.Parity.String()
}

type countChecker struct {
	Checker numberChecker
	Count   int
}

func (c countChecker) Validate(code domain.Code) bool {
	return lo.CountBy(code[:], c.Checker.Check) == c.Count
}

func (c countChecker) String() string {
	return fmt.Sprintf("have %d of %v", c.Count, c.Checker)
}

type repetitionCounter struct {
	Result int
}

func (c repetitionCounter) Validate(code domain.Code) bool {
	grouped := lo.GroupBy(code[:], func(it int) int {
		return it
	})

	maxCount := len(lo.MaxBy(lo.Values(grouped), func(ints []int, ints2 []int) bool {
		return len(ints) > len(ints2)
	}))
	repCount := maxCount - 1

	return repCount == c.Result
}

func (c repetitionCounter) String() string {
	if c.Result == 0 {
		return "no repetitions"
	}

	return fmt.Sprintf("has %d doubles", c.Result+1)
}
