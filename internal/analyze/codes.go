package analyze

import (
	"sort"
	"turing-machine/internal/domain"

	"github.com/samber/lo"
)

func Codes(codes []domain.Code) []NumberStats {
	result := make([]NumberStats, 0, domain.CodeLength)
	for i := 0; i < domain.CodeLength; i++ {
		seen := make(map[int]struct{})
		for _, code := range codes {
			seen[code[i]] = struct{}{}
		}

		values := lo.Keys(seen)
		sort.Ints(values)

		result = append(result, NumberStats{
			Item:   domain.CodeItem(i),
			Values: values,
		})
	}

	return result
}
