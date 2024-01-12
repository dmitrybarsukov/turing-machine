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

// ItemComparedToConst
// Число (▲,■,●) в сравнении с (1,2,3,4,5)
func ItemComparedToConst(item domain.CodeItem, constant int) []domain.Validator {
	variants := []Compare{Less, Equal, More}
	return util.Combine1(variants, func(compare Compare) domain.Validator {
		return constComparator{
			Item:    item,
			Const:   constant,
			Compare: compare,
		}
	})
}

// ItemComparedToOtherItem
// Число (▲,■,●) в сравнении с числом (▲,■,●)
func ItemComparedToOtherItem(item1, item2 domain.CodeItem) []domain.Validator {
	variants := []Compare{Less, Equal, More}
	return util.Combine1(variants, func(compare Compare) domain.Validator {
		return twoItemComparator{
			Item1:   item1,
			Item2:   item2,
			Compare: compare,
		}
	})
}

// ItemsSumComparedToConst
// Сумма чисел (▲,■,●) в сравнении с числом (1,2...15)
func ItemsSumComparedToConst(items []domain.CodeItem, sum int) []domain.Validator {
	variants := []Compare{Less, Equal, More}
	return util.Combine1(variants, func(compare Compare) domain.Validator {
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

// AnyItemsPairCompared
// Сравнение чисел в одной из пар (▲,■), (▲,●), (■,●)
func AnyItemsPairCompared() []domain.Validator {
	return lo.Flatten([][]domain.Validator{
		ItemComparedToOtherItem(domain.CodeItemTriangle, domain.CodeItemSquare),
		ItemComparedToOtherItem(domain.CodeItemTriangle, domain.CodeItemCircle),
		ItemComparedToOtherItem(domain.CodeItemSquare, domain.CodeItemCircle),
	})
}

// CountOfNumber
// Количество чисел (1,2,3,4,5) в коде
func CountOfNumber(number int) []domain.Validator {
	variants := []int{0, 1, 2, 3}
	return util.Combine1(variants, func(count int) domain.Validator {
		return countChecker{
			Checker: equalityChecker{Number: number},
			Count:   count,
		}
	})
}

// CountOfParity
// Количество (чётных,нечётных) чисел в коде
func CountOfParity(parity Parity) []domain.Validator {
	variants := []int{0, 1, 2, 3}
	return util.Combine1(variants, func(count int) domain.Validator {
		return countChecker{
			Checker: parityChecker{Parity: parity},
			Count:   count,
		}
	})
}

// HasMajorParity
// Количество чётных чисел по сравнению с количеством нечётных
func HasMajorParity() []domain.Validator {
	variants := []Parity{Even, Odd}
	return util.Combine1(variants, func(parity Parity) domain.Validator {
		return majorParityChecker{
			Parity: parity,
		}
	})
}

// CountOfRepetitions
// Имеет (0,1,2) повторяющихся цифр
func CountOfRepetitions() []domain.Validator {
	variants := []int{0, 1, 2}
	return util.Combine1(variants, func(count int) domain.Validator {
		return repetitionCounter{
			Result: count,
		}
	})
}

// PairOfNumbersExist
// В коде (есть,нет) пара одинаковых чисел
func PairOfNumbersExist() []domain.Validator {
	variants := []bool{false, true}
	return util.Combine1(variants, func(result bool) domain.Validator {
		return hasSameNumbersPairChecker{
			Result: result,
		}
	})
}

// CountOfAnyNumber
// Количество каких то из чисел (1,2,3,4,5) в коде
func CountOfAnyNumber(values ...int) []domain.Validator {
	return lo.FlatMap(values, func(it int, _ int) []domain.Validator {
		return CountOfNumber(it)
	})
}

// ItemIsGreatest
// Одно из чисел (▲,■,●) больше остальных
func ItemIsGreatest() []domain.Validator {
	return util.Combine1(codeItemVariants, func(item domain.CodeItem) domain.Validator {
		return itemIsOutstandingChecker{
			Item:    item,
			Compare: More,
		}
	})
}

// ItemIsLeast
// Одно из чисел (▲,■,●) больше остальных
func ItemIsLeast() []domain.Validator {
	return util.Combine1(codeItemVariants, func(item domain.CodeItem) domain.Validator {
		return itemIsOutstandingChecker{
			Item:    item,
			Compare: Less,
		}
	})
}

// ItemIsOutstanding
// Одно из чисел (▲,■,●) больше или меньше остальных
func ItemIsOutstanding() []domain.Validator {
	compareVariants := []Compare{Less, More}
	return util.Combine2(codeItemVariants, compareVariants, func(item domain.CodeItem, compare Compare) domain.Validator {
		return itemIsOutstandingChecker{
			Item:    item,
			Compare: compare,
		}
	})
}

// CodeHasOrder
// Код отсортирован по возрастанию, либо по убыванию, либо никак
func CodeHasOrder() []domain.Validator {
	variants := []Order{Ascending, None, Descending}
	return util.Combine1(variants, func(order Order) domain.Validator {
		return orderStrictChecker{
			Order: order,
		}
	})
}

// ItemHasParity
// Указанное число чётное либо нечётное
func ItemHasParity(item domain.CodeItem) []domain.Validator {
	variants := []Parity{Even, Odd}
	return util.Combine1(variants, func(parity Parity) domain.Validator {
		return itemParityChecker{
			Item:   item,
			Parity: parity,
		}
	})
}

// SumHasParity
// Сумма цифр чётная либо нечётная
func SumHasParity() []domain.Validator {
	variants := []Parity{Even, Odd}
	return util.Combine1(variants, func(parity Parity) domain.Validator {
		return sumParityChecker{
			Parity: parity,
		}
	})
}

// HasSequence
// / В коде (есть,нет) пара или больше последовательных чисел в указанном порядке
func HasSequence(order Order) []domain.Validator {
	variants := []bool{false, true}
	return util.Combine1(variants, func(result bool) domain.Validator {
		return hasSequenceChecker{
			Order:  order,
			Result: result,
		}
	})
}

// HasAnySequence
// В коде (есть,нет) пара или больше последовательных чисел в любом порядке
func HasAnySequence() []domain.Validator {
	variants := []bool{false, true}
	return util.Combine1(variants, func(result bool) domain.Validator {
		return hasAnySequenceChecker{
			Result: result,
		}
	})
}
