package parser

import (
	"errors"
	"fmt"
	"strings"
	"turing-machine/internal/domain"
	"turing-machine/internal/domain/validator"

	"github.com/samber/lo"
)

var mappingCodeItem = map[string]domain.CodeItem{
	"triangle": domain.CodeItemTriangle,
	"square":   domain.CodeItemSquare,
	"circle":   domain.CodeItemCircle,
}
var allCodeItem = lo.Values(mappingCodeItem)

var mappingCompare = map[string]validator.Compare{
	"less":  validator.Less,
	"equal": validator.Equal,
	"more":  validator.More,
}
var allCompare = lo.Values(mappingCompare)

var mappingOrder = map[string]validator.Order{
	"asc":  validator.Ascending,
	"none": validator.None,
	"desc": validator.Descending,
}
var allOrder = lo.Values(mappingOrder)

var mappingParity = map[string]validator.Parity{
	"even": validator.Even,
	"odd":  validator.Odd,
}
var allParity = lo.Values(mappingParity)

var mappingCount = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
}
var allCount = lo.Values(mappingCount)

var mappingBool = map[string]bool{
	"false": false,
	"true":  true,
}
var allBool = lo.Values(mappingBool)

func parseEnum[T comparable](value string, mapping map[string]T) (T, error) {
	item, ok := mapping[value]
	if !ok {
		var t T
		return t, fmt.Errorf("unknown %T: %s", t, value)
	}

	return item, nil
}

func parseVariants[T comparable](clues []string, mapping map[string]T) ([]T, error) {
	values := lo.SliceToMap(lo.Values(mapping), func(it T) (T, struct{}) {
		return it, struct{}{}
	})

	for _, option := range clues {
		optionNorm, exclude := parseNot(option)
		item, err := parseEnum(optionNorm, mapping)
		if err != nil {
			return nil, err
		}

		if !exclude {
			values = map[T]struct{}{item: {}}
		} else {
			delete(values, item)
		}
	}

	if len(values) == 0 {
		return nil, errors.New("all variants excluded")
	}

	return lo.Keys(values), nil
}

func parseNot(value string) (string, bool) {
	const prefix = "not "

	if strings.HasPrefix(value, prefix) {
		return strings.TrimPrefix(value, prefix), true
	}

	return value, false
}
