package domain

import (
	"strconv"
	"strings"

	"github.com/dmitrybarsukov/turing-machine/internal/util"

	"github.com/samber/lo"
)

type CodeItem int

const (
	CodeItem0 CodeItem = iota
	CodeItem1
	CodeItem2

	CodeItemTriangle = CodeItem0
	CodeItemSquare   = CodeItem1
	CodeItemCircle   = CodeItem2
)

func (c CodeItem) Index() int {
	return int(c)
}

func (c CodeItem) String() string {
	switch c {
	case CodeItemTriangle:
		return "▲"
	case CodeItemSquare:
		return "■"
	case CodeItemCircle:
		return "●"
	default:
		return "?"
	}
}

const CodeLength = 3

type Code [CodeLength]int

func (c Code) Get(item CodeItem) int {
	return c[item.Index()]
}

func (c Code) String() string {
	var builder strings.Builder
	for _, i := range c {
		builder.WriteString(strconv.Itoa(i))
	}

	return builder.String()
}

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
