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

var mappingCompare = map[string]validator.Compare{
	"less":  validator.Less,
	"equal": validator.Equal,
	"more":  validator.More,
}

var mappingOrder = map[string]validator.Order{
	"asc":  validator.Ascending,
	"none": validator.None,
	"desc": validator.Descending,
}

var mappingParity = map[string]validator.Parity{
	"even": validator.Even,
	"odd":  validator.Odd,
}

var mappingCount = map[string]validator.Count{
	"0": validator.Zero,
	"1": validator.One,
	"2": validator.Two,
	"3": validator.Three,
}

var mappingBool = map[string]bool{
	"false": false,
	"true":  true,
}

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
