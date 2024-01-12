package validator

import (
	"testing"
	"turing-machine/internal/domain"
)

func TestConstComparator(t *testing.T) {
	t.Run("[0] < 2", func(t *testing.T) {
		tt := newTester(t, constComparator{
			Item:    domain.CodeItem0,
			Const:   2,
			Compare: Less,
		})

		tt.test(155, true)
		tt.test(244, false)
		tt.test(333, false)
		tt.test(422, false)
		tt.test(511, false)
	})

	t.Run("[1] = 3", func(t *testing.T) {
		tt := newTester(t, constComparator{
			Item:    domain.CodeItem1,
			Const:   3,
			Compare: Equal,
		})

		tt.test(155, false)
		tt.test(244, false)
		tt.test(333, true)
		tt.test(422, false)
		tt.test(511, false)
	})

	t.Run("[2] = 4", func(t *testing.T) {
		tt := newTester(t, constComparator{
			Item:    domain.CodeItem2,
			Const:   4,
			Compare: More,
		})

		tt.test(155, true)
		tt.test(244, false)
		tt.test(333, false)
		tt.test(422, false)
		tt.test(511, false)
	})
}

func TestTwoItemComparator(t *testing.T) {
	t.Run("[0] < [1]", func(t *testing.T) {
		tt := newTester(t, twoItemComparator{
			Item1:   domain.CodeItem0,
			Item2:   domain.CodeItem1,
			Compare: Less,
		})

		tt.test(155, true)
		tt.test(244, true)
		tt.test(333, false)
		tt.test(422, false)
		tt.test(511, false)
	})

	t.Run("[1] = [2]", func(t *testing.T) {
		tt := newTester(t, twoItemComparator{
			Item1:   domain.CodeItem1,
			Item2:   domain.CodeItem2,
			Compare: Equal,
		})

		tt.test(151, false)
		tt.test(243, false)
		tt.test(333, true)
		tt.test(422, true)
		tt.test(512, false)
	})

	t.Run("[0] > [2]", func(t *testing.T) {
		tt := newTester(t, twoItemComparator{
			Item1:   domain.CodeItem0,
			Item2:   domain.CodeItem2,
			Compare: More,
		})

		tt.test(155, false)
		tt.test(244, false)
		tt.test(333, false)
		tt.test(422, true)
		tt.test(511, true)
	})
}

func TestItemsSumComparator(t *testing.T) {
	t.Run("[0] + [1] < 6", func(t *testing.T) {
		tt := newTester(t, itemsSumComparator{
			Items:   [domain.CodeLength]bool{true, true, false},
			Sum:     6,
			Compare: Less,
		})

		tt.test(155, false)
		tt.test(214, true)
		tt.test(113, true)
		tt.test(543, false)
		tt.test(333, false)
	})

	t.Run("[0] + [1] + [2] > 10", func(t *testing.T) {
		tt := newTester(t, itemsSumComparator{
			Items:   [domain.CodeLength]bool{true, true, true},
			Sum:     10,
			Compare: More,
		})

		tt.test(155, true)
		tt.test(214, false)
		tt.test(113, false)
		tt.test(543, true)
		tt.test(333, false)
	})
}
