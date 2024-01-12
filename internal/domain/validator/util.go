package validator

import (
	"turing-machine/internal/domain"

	"github.com/samber/lo"
)

type validatorWithValue[T any] interface {
	domain.Validator
	WithValue(value T) domain.Validator
}

func makeValidators[T any](
	base validatorWithValue[T],
	variants []T,
) []domain.Validator {
	return lo.Map(variants, func(it T, _ int) domain.Validator {
		return base.WithValue(it)
	})
}

func compare(value1, value2 int) Compare {
	if value1 < value2 {
		return Less
	} else if value1 > value2 {
		return More
	} else {
		return Equal
	}
}

func getOrder(arr []int) Order {
	var isAsc, isDesc = true, true
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] <= arr[i+1] {
			isDesc = false
		}

		if arr[i] >= arr[i+1] {
			isAsc = false
		}
	}

	if isAsc {
		return Ascending
	}

	if isDesc {
		return Descending
	}

	return Unordered
}

func getParity(value int) Parity {
	return []Parity{Even, Odd}[value%2]
}
