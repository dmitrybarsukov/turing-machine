package validator

import (
	"turing-machine/internal/domain"
	"turing-machine/internal/util"

	"github.com/samber/lo"
)

var codeItemVariants = []domain.CodeItem{
	domain.CodeItemTriangle,
	domain.CodeItemSquare,
	domain.CodeItemCircle,
}

func ItemComparedToConst(item domain.CodeItem, constant int) []domain.Validator {
	variants := []Compare{Less, Equal, More}
	return util.Map(variants, func(compare Compare) domain.Validator {
		return constComparator{
			Item:    item,
			Const:   constant,
			Compare: compare,
		}
	})
}

func ItemComparedToOtherItem(item1, item2 domain.CodeItem) []domain.Validator {
	variants := []Compare{Less, Equal, More}
	return util.Map(variants, func(compare Compare) domain.Validator {
		return twoItemComparator{
			Item1:   item1,
			Item2:   item2,
			Compare: compare,
		}
	})
}

func ItemsSumComparedToConst(items []domain.CodeItem, sum int) []domain.Validator {
	variants := []Compare{Less, Equal, More}
	return util.Map(variants, func(compare Compare) domain.Validator {
		return itemsSumComparator{
			Items: [3]bool{
				lo.Contains(items, domain.CodeItem0),
				lo.Contains(items, domain.CodeItem1),
				lo.Contains(items, domain.CodeItem2),
			},
			Sum:     sum,
			Compare: compare,
		}
	})
}

func AnyItemsPairCompared() []domain.Validator {
	return lo.Flatten([][]domain.Validator{
		ItemComparedToOtherItem(domain.CodeItemTriangle, domain.CodeItemSquare),
		ItemComparedToOtherItem(domain.CodeItemTriangle, domain.CodeItemCircle),
		ItemComparedToOtherItem(domain.CodeItemSquare, domain.CodeItemCircle),
	})
}

func CountOfNumber(number int) []domain.Validator {
	variants := []int{0, 1, 2, 3}
	return util.Map(variants, func(count int) domain.Validator {
		return countChecker{
			Checker: equalityChecker{Number: number},
			Count:   count,
		}
	})
}

func CountOfParity(parity Parity) []domain.Validator {
	variants := []int{0, 1, 2, 3}
	return util.Map(variants, func(count int) domain.Validator {
		return countChecker{
			Checker: parityChecker{Parity: parity},
			Count:   count,
		}
	})
}

func HasMajorParity() []domain.Validator {
	variants := []Parity{Even, Odd}
	return util.Map(variants, func(parity Parity) domain.Validator {
		return majorParityChecker{
			Parity: parity,
		}
	})
}

func CountOfRepetitions() []domain.Validator {
	variants := []int{0, 1, 2}
	return util.Map(variants, func(count int) domain.Validator {
		return repetitionCounter{
			Result: count,
		}
	})
}

func PairOfNumbersExist() []domain.Validator {
	variants := []bool{false, true}
	return util.Map(variants, func(result bool) domain.Validator {
		return hasSameNumbersPairChecker{
			Result: result,
		}
	})
}

func CountOfAnyNumber(values ...int) []domain.Validator {
	return lo.FlatMap(values, func(it int, _ int) []domain.Validator {
		return CountOfNumber(it)
	})
}

func ItemIsGreatest() []domain.Validator {
	return util.Map(codeItemVariants, func(item domain.CodeItem) domain.Validator {
		return itemIsOutstandingChecker{
			Item:    item,
			Compare: More,
		}
	})
}

func ItemIsLeast() []domain.Validator {
	return util.Map(codeItemVariants, func(item domain.CodeItem) domain.Validator {
		return itemIsOutstandingChecker{
			Item:    item,
			Compare: Less,
		}
	})
}

func ItemIsOutstanding() []domain.Validator {
	compareVariants := []Compare{Less, More}
	return util.CrossMap(codeItemVariants, compareVariants, func(item domain.CodeItem, compare Compare) domain.Validator {
		return itemIsOutstandingChecker{
			Item:    item,
			Compare: compare,
		}
	})
}

func CodeHasOrder() []domain.Validator {
	variants := []Order{Ascending, None, Descending}
	return util.Map(variants, func(order Order) domain.Validator {
		return orderStrictChecker{
			Order: order,
		}
	})
}

func ItemHasParity(item domain.CodeItem) []domain.Validator {
	variants := []Parity{Even, Odd}
	return util.Map(variants, func(parity Parity) domain.Validator {
		return itemParityChecker{
			Item:   item,
			Parity: parity,
		}
	})
}

func SumHasParity() []domain.Validator {
	variants := []Parity{Even, Odd}
	return util.Map(variants, func(parity Parity) domain.Validator {
		return sumParityChecker{
			Parity: parity,
		}
	})
}

func AnyItemParity() []domain.Validator {
	variants := []Parity{Even, Odd}
	return util.CrossMap(codeItemVariants, variants, func(item domain.CodeItem, parity Parity) domain.Validator {
		return itemParityChecker{
			Item:   item,
			Parity: parity,
		}
	})
}

func HasSequence(order Order) []domain.Validator {
	variants := []bool{false, true}
	return util.Map(variants, func(result bool) domain.Validator {
		return hasSequenceChecker{
			Order:  order,
			Result: result,
		}
	})
}

func HasAnySequence() []domain.Validator {
	variants := []bool{false, true}
	return util.Map(variants, func(result bool) domain.Validator {
		return hasAnySequenceChecker{
			Result: result,
		}
	})
}

func AnyItemComparedToConst(compare Compare, constant int) []domain.Validator {
	return util.Map(codeItemVariants, func(item domain.CodeItem) domain.Validator {
		return constComparator{
			Item:    item,
			Const:   constant,
			Compare: compare,
		}
	})
}
