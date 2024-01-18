package validator

import (
	"fmt"
	"strings"

	"github.com/dmitrybarsukov/turing-machine/internal/domain"

	"github.com/samber/lo"
)

type constComparator struct {
	Item    domain.CodeItem
	Const   int
	Compare Compare
}

func (c constComparator) Validate(code domain.Code) bool {
	return compare(code.Get(c.Item), c.Const) == c.Compare
}

func (c constComparator) String() string {
	return fmt.Sprintf("%v %v %d", c.Item, c.Compare, c.Const)
}

type twoItemComparator struct {
	Item1   domain.CodeItem
	Item2   domain.CodeItem
	Compare Compare
}

func (c twoItemComparator) Validate(code domain.Code) bool {
	return compare(code.Get(c.Item1), code.Get(c.Item2)) == c.Compare
}

func (c twoItemComparator) String() string {
	return fmt.Sprintf("%v %v %v", c.Item1, c.Compare, c.Item2)
}

type itemsSumComparator struct {
	Items   [domain.CodeLength]bool
	Sum     int
	Compare Compare
}

func (c itemsSumComparator) Validate(code domain.Code) bool {
	sum := 0
	for idx, ok := range c.Items {
		if ok {
			sum += code[idx]
		}
	}
	return compare(sum, c.Sum) == c.Compare
}

func (c itemsSumComparator) String() string {
	itemStrs := lo.FilterMap(c.Items[:], func(ok bool, idx int) (string, bool) {
		return domain.CodeItem(idx).String(), ok
	})
	return fmt.Sprintf("%s %v %v", strings.Join(itemStrs, " + "), c.Compare, c.Sum)
}
