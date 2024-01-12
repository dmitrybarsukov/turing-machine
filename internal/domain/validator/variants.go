package validator

import (
	"turing-machine/internal/domain"
)

var codeItemVariants = []domain.CodeItem{
	domain.CodeItemTriangle,
	domain.CodeItemSquare,
	domain.CodeItemCircle,
}

var compareVariants = []Compare{
	Less,
	Equal,
	More,
}

var differentCompareVariants = []Compare{
	Less,
	More,
}

var orderVariants = []Order{
	Ascending,
	None,
	Descending,
}

var parityVariants = []Parity{
	Even,
	Odd,
}

var boolVariants = []bool{
	false,
	true,
}

var countVariants = []int{
	0, 1, 2, 3,
}

var repetitionVariants = []int{
	0, 1, 2,
}
