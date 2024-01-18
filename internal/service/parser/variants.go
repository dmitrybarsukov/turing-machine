package parser

import (
	"fmt"

	"github.com/dmitrybarsukov/turing-machine/internal/domain"
	"github.com/dmitrybarsukov/turing-machine/internal/service/validator"
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

var mappingCompare = map[string]validator.Compare{
	"less":  validator.Less,
	"equal": validator.Equal,
	"more":  validator.More,
}

func parseEnum[T comparable](value string, mapping map[string]T) (T, error) {
	item, ok := mapping[value]
	if !ok {
		var t T
		return t, fmt.Errorf("unknown %T: %s", t, value)
	}

	return item, nil
}
