package parser

import (
	"fmt"
	"turing-machine/internal/domain"
	"turing-machine/internal/domain/validator"
)

var mappingCodeItem = map[string]domain.CodeItem{
	"triangle": domain.CodeItemTriangle,
	"square":   domain.CodeItemSquare,
	"circle":   domain.CodeItemCircle,
}

var mappingParity = map[string]validator.Parity{
	"even": validator.Even,
	"odd":  validator.Odd,
}

func parseEnum[T comparable](value string, mapping map[string]T) (T, error) {
	item, ok := mapping[value]
	if !ok {
		var t T
		return t, fmt.Errorf("unknown %T: %s", t, value)
	}

	return item, nil
}
