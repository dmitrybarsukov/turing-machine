package parser

import (
	"turing-machine/internal/domain"
	"turing-machine/internal/domain/turing"

	"github.com/samber/lo"
)

type ParseResult struct {
	HyperMachine turing.HyperMachine
	Tests        []Test
}

type Test struct {
	Code      domain.Code
	Validator rune
	Result    bool
}

func (t Test) FilterMachines(machines []turing.Machine) []turing.Machine {
	if t.Result {
		return lo.Filter(machines, func(it turing.Machine, _ int) bool {
			validator, ok := it.Validators[t.Validator]
			if !ok {
				return true
			}
			return validator.Validate(t.Code)
		})
	}

	return lo.Filter(machines, func(it turing.Machine, _ int) bool {
		validator, ok := it.Validators[t.Validator]
		if !ok {
			return true
		}
		return !validator.Validate(t.Code)
	})
}
