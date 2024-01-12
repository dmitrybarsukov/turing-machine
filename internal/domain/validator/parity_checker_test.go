package validator

import (
	"testing"
	"turing-machine/internal/domain"
)

func TestParityChecker(t *testing.T) {
	t.Run("[0] is even", func(t *testing.T) {
		tt := newTester(t, itemParityChecker{Item: domain.CodeItem0, result: Even})

		tt.test(155, false)
		tt.test(322, false)
		tt.test(333, false)
		tt.test(433, true)
		tt.test(455, true)
	})

	t.Run("[2] is odd", func(t *testing.T) {
		tt := newTester(t, itemParityChecker{Item: domain.CodeItem2, result: Odd})

		tt.test(155, true)
		tt.test(322, false)
		tt.test(333, true)
		tt.test(433, true)
		tt.test(452, false)
	})
}
