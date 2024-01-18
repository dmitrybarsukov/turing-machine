package turing

import (
	"sort"

	"github.com/dmitrybarsukov/turing-machine/internal/domain"
	"github.com/dmitrybarsukov/turing-machine/internal/util"

	"github.com/samber/lo"
)

type HyperMachine struct {
	Validators map[rune][]domain.Validator
	Codes      []domain.Code
}

func (m HyperMachine) GetAllMachines() []Machine {
	keys := lo.Keys(m.Validators)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	var validatorCombinations [][]domain.Validator
	for _, key := range keys {
		validators := m.Validators[key]
		validatorCombinations = util.Cross(validatorCombinations, validators)
	}

	return lo.Map(validatorCombinations, func(it []domain.Validator, _ int) Machine {
		validators := make(map[rune]domain.Validator, len(keys))
		for i, key := range keys {
			validators[key] = it[i]
		}

		return Machine{
			Validators: validators,
			Codes:      m.Codes,
		}
	})
}
