package domain

import (
	"turing-machine/internal/util"

	"github.com/samber/lo"
)

var AllCodes []Code

func init() {
	digits := []int{1, 2, 3, 4, 5}
	allCombinations := util.Combinations(digits, CodeLength)
	AllCodes = lo.Map(allCombinations, func(it []int, _ int) Code {
		var code Code
		copy(code[:], it)
		return code
	})
}
