package analyze

import (
	"fmt"
	"strconv"
	"strings"
	"turing-machine/internal/domain"

	"github.com/samber/lo"
)

type NumberStats struct {
	Item   domain.CodeItem
	Values []int
}

func (n NumberStats) String() string {
	values := lo.Map(n.Values, func(it int, _ int) string {
		return strconv.Itoa(it)
	})
	return fmt.Sprintf("%v: [%s]", n.Item, strings.Join(values, ", "))
}

type ValidatorStats struct {
	Key        rune
	Validator  domain.Validator
	Confidence float64
}

func (v ValidatorStats) String() string {
	return fmt.Sprintf("%c: [%v] - %.0f %%", v.Key, v.Validator, v.Confidence*100)
}
