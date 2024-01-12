package parser

import (
	"turing-machine/internal/domain"
	turing2 "turing-machine/internal/service/turing"

	"github.com/samber/lo"
)

type ParseResult struct {
	HyperMachine turing2.HyperMachine
	Tests        []Test
}

type Test struct {
	Code      domain.Code
	Validator rune
	Result    bool
}

func (t Test) FilterMachines(machines []turing2.Machine) []turing2.Machine {
	if t.Result {
		return lo.Filter(machines, func(it turing2.Machine, _ int) bool {
			validator, ok := it.Validators[t.Validator]
			if !ok {
				return true
			}
			return validator.Validate(t.Code)
		})
	}

	return lo.Filter(machines, func(it turing2.Machine, _ int) bool {
		validator, ok := it.Validators[t.Validator]
		if !ok {
			return true
		}
		return !validator.Validate(t.Code)
	})
}
