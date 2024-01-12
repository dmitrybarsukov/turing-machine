package validator

import (
	"testing"
	"turing-machine/internal/domain"
)

func TestMajorParityChecker(t *testing.T) {
	t.Run("has more even than odd", func(t *testing.T) {
		tt := newTester(t, majorParityChecker{Parity: Even})

		tt.test(155, false)
		tt.test(322, true)
		tt.test(333, false)
		tt.test(433, false)
		tt.test(512, false)
	})

	t.Run("has more odd than even", func(t *testing.T) {
		tt := newTester(t, majorParityChecker{Parity: Odd})

		tt.test(521, true)
		tt.test(522, false)
		tt.test(123, true)
		tt.test(251, true)
		tt.test(241, false)
	})
}

func TestHasSameNumbersPairChecker(t *testing.T) {
	t.Run("has no same number pairs", func(t *testing.T) {
		tt := newTester(t, hasSameNumbersPairChecker{Result: false})

		tt.test(155, false)
		tt.test(322, false)
		tt.test(333, true)
		tt.test(433, false)
		tt.test(512, true)
	})

	t.Run("has some pairs", func(t *testing.T) {
		tt := newTester(t, hasSameNumbersPairChecker{Result: true})

		tt.test(155, true)
		tt.test(232, true)
		tt.test(333, false)
		tt.test(112, true)
		tt.test(512, false)
	})
}

func TestItemIsOutstandingChecker(t *testing.T) {
	t.Run("[0] is less", func(t *testing.T) {
		tt := newTester(t, itemIsOutstandingChecker{Compare: Less, Item: domain.CodeItem0})

		tt.test(155, true)
		tt.test(322, false)
		tt.test(333, false)
		tt.test(433, false)
		tt.test(455, true)
	})

	t.Run("[1] is less", func(t *testing.T) {
		tt := newTester(t, itemIsOutstandingChecker{Compare: Less, Item: domain.CodeItem1})

		tt.test(435, true)
		tt.test(322, false)
		tt.test(333, false)
		tt.test(432, false)
		tt.test(415, true)
	})

	t.Run("[2] is more", func(t *testing.T) {
		tt := newTester(t, itemIsOutstandingChecker{Compare: More, Item: domain.CodeItem2})

		tt.test(435, true)
		tt.test(322, false)
		tt.test(333, false)
		tt.test(432, false)
		tt.test(415, true)
	})
}

func TestOrderStrictChecker(t *testing.T) {
	t.Run("Asc", func(t *testing.T) {
		tt := newTester(t, orderStrictChecker{Order: Ascending})

		tt.test(435, false)
		tt.test(322, false)
		tt.test(333, false)
		tt.test(432, false)
		tt.test(345, true)
		tt.test(355, false)
	})

	t.Run("Desc", func(t *testing.T) {
		tt := newTester(t, orderStrictChecker{Order: Descending})

		tt.test(435, false)
		tt.test(321, true)
		tt.test(333, false)
		tt.test(432, true)
		tt.test(345, false)
		tt.test(522, false)
	})

	t.Run("None", func(t *testing.T) {
		tt := newTester(t, orderStrictChecker{Order: None})

		tt.test(435, true)
		tt.test(321, false)
		tt.test(333, true)
		tt.test(432, false)
		tt.test(345, false)
		tt.test(522, true)
	})
}
