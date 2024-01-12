package validator

import (
	"testing"
	"turing-machine/internal/domain"
)

func TestItemParityChecker(t *testing.T) {
	t.Run("[0] is even", func(t *testing.T) {
		tt := newTester(t, itemParityChecker{Item: domain.CodeItem0, Parity: Even})

		tt.test(155, false)
		tt.test(322, false)
		tt.test(333, false)
		tt.test(433, true)
		tt.test(455, true)
	})

	t.Run("[2] is odd", func(t *testing.T) {
		tt := newTester(t, itemParityChecker{Item: domain.CodeItem2, Parity: Odd})

		tt.test(155, true)
		tt.test(322, false)
		tt.test(333, true)
		tt.test(433, true)
		tt.test(452, false)
	})
}

func TestSumParityChecker(t *testing.T) {
	t.Run("sum is even", func(t *testing.T) {
		tt := newTester(t, sumParityChecker{Parity: Even})

		tt.test(155, false)
		tt.test(323, true)
		tt.test(331, false)
		tt.test(433, true)
		tt.test(455, true)
	})

	t.Run("sum is odd", func(t *testing.T) {
		tt := newTester(t, sumParityChecker{Parity: Odd})

		tt.test(555, true)
		tt.test(322, true)
		tt.test(333, true)
		tt.test(433, false)
		tt.test(451, false)
	})
}
