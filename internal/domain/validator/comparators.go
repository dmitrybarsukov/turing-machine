package validator

import (
	"fmt"
	"strings"
	"turing-machine/internal/domain"

	"github.com/samber/lo"
)

type constComparator struct {
	Item  domain.CodeItem
	Const int

	result Compare
}

func (c constComparator) Validate(code domain.Code) bool {
	return compare(code.Get(c.Item), c.Const) == c.result
}

func (c constComparator) String() string {
	return fmt.Sprintf("%v %v %d", c.Item, c.result, c.Const)
}

func (c constComparator) WithValue(value Compare) domain.Validator {
	c.result = value
	return c
}

func ItemComparedToConst(item domain.CodeItem, constant int) []domain.Validator {
	return makeValidators[Compare](constComparator{Item: item, Const: constant}, compareVariants)
}

type itemComparator struct {
	Item1 domain.CodeItem
	Item2 domain.CodeItem

	result Compare
}

func (c itemComparator) Validate(code domain.Code) bool {
	return compare(code.Get(c.Item1), code.Get(c.Item2)) == c.result
}

func (c itemComparator) String() string {
	return fmt.Sprintf("%v %v %v", c.Item1, c.result, c.Item2)
}

func (c itemComparator) WithValue(value Compare) domain.Validator {
	c.result = value
	return c
}

func ItemComparedToOtherItem(item1, item2 domain.CodeItem) []domain.Validator {
	return makeValidators[Compare](itemComparator{Item1: item1, Item2: item2}, compareVariants)
}

type itemsSumComparator struct {
	Items []domain.CodeItem
	Sum   int

	result Compare
}

func (c itemsSumComparator) Validate(code domain.Code) bool {
	sum := lo.SumBy[domain.CodeItem, int](c.Items, func(item domain.CodeItem) int {
		return code.Get(item)
	})
	return compare(sum, c.Sum) == c.result
}

func (c itemsSumComparator) String() string {
	itemStrs := lo.Map(c.Items, func(it domain.CodeItem, _ int) string {
		return it.String()
	})
	return fmt.Sprintf("%s %v %v", strings.Join(itemStrs, " + "), c.result, c.Sum)
}

func (c itemsSumComparator) WithValue(value Compare) domain.Validator {
	c.result = value
	return c
}

func ItemsSumComparedToConst(items []domain.CodeItem, sum int) []domain.Validator {
	return makeValidators[Compare](itemsSumComparator{Items: items, Sum: sum}, compareVariants)
}

func ItemsMultiComparable() []domain.Validator {
	return lo.Flatten([][]domain.Validator{
		ItemComparedToOtherItem(domain.CodeItemTriangle, domain.CodeItemSquare),
		ItemComparedToOtherItem(domain.CodeItemSquare, domain.CodeItemCircle),
		ItemComparedToOtherItem(domain.CodeItemCircle, domain.CodeItemTriangle),
	})
}
