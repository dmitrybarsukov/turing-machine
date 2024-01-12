package domain

import (
	"fmt"
	"strconv"
	"strings"
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

const codeLength = 3

type Code [codeLength]int

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

type Validator interface {
	Validate(code Code) bool
	fmt.Stringer
}
