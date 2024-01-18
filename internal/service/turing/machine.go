package turing

import (
	"fmt"
	"sort"
	"strings"

	"github.com/dmitrybarsukov/turing-machine/internal/domain"

	"github.com/samber/lo"
)

type Machine struct {
	Validators map[rune]domain.Validator
	Codes      []domain.Code
}

func (m Machine) String() string {
	keys := lo.Keys(m.Validators)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	items := make([]string, 0, len(m.Validators))
	for _, key := range keys {
		validator := m.Validators[key]
		items = append(items, fmt.Sprintf("%c: [%v]", key, validator))
	}

	var suffix string

	codes := m.getGoodCodes()
	if len(codes) == 1 {
		suffix = "Solution: " + codes[0].String()
	} else {
		suffix = fmt.Sprintf("Solutions count: %d", len(codes))
	}

	return "Machine{ " + strings.Join(items, ", ") + " | " + suffix + " }"
}

func (m Machine) getGoodCodes() []domain.Code {
	return lo.Filter(m.Codes, func(code domain.Code, _ int) bool {
		return lo.EveryBy(lo.Values(m.Validators), func(val domain.Validator) bool {
			return val.Validate(code)
		})
	})
}

func (m Machine) HasSolution() bool {
	return m.Solution() != domain.Code{}
}

func (m Machine) Solution() domain.Code {
	codes := m.getGoodCodes()
	if len(codes) == 1 {
		return codes[0]
	}

	return domain.Code{}
}
