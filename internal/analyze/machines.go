package analyze

import (
	"sort"
	"turing-machine/internal/domain"
	"turing-machine/internal/service/turing"

	"github.com/samber/lo"
)

func Validators(machines []turing.Machine) []ValidatorStats {
	allKeysMap := make(map[rune]struct{})
	for _, m := range machines {
		for k := range m.Validators {
			allKeysMap[k] = struct{}{}
		}
	}

	allKeys := lo.Keys(allKeysMap)
	sort.Slice(allKeys, func(i, j int) bool {
		return allKeys[i] < allKeys[j]
	})

	result := make([]ValidatorStats, 0, len(allKeys))
	for _, key := range allKeys {
		counts := make(map[domain.Validator]int)
		maxCount := 0
		maxValidator := domain.Validator(nil)
		for _, m := range machines {
			validator := m.Validators[key]
			count := counts[validator] + 1
			counts[validator] = count
			if count > maxCount {
				maxCount = count
				maxValidator = validator
			}
		}

		result = append(result, ValidatorStats{
			Key:        key,
			Validator:  maxValidator,
			Confidence: float64(maxCount) / float64(len(machines)),
		})
	}
	return result
}
