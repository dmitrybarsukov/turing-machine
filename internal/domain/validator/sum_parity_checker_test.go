package validator

import (
	"testing"
)

func TestSumParityChecker(t *testing.T) {
	t.Run("sum is even", func(t *testing.T) {
		tt := newTester(t, sumParityChecker{result: Even})

		tt.test(155, false)
		tt.test(323, true)
		tt.test(331, false)
		tt.test(433, true)
		tt.test(455, true)
	})

	t.Run("sum is odd", func(t *testing.T) {
		tt := newTester(t, sumParityChecker{result: Odd})

		tt.test(555, true)
		tt.test(322, true)
		tt.test(333, true)
		tt.test(433, false)
		tt.test(451, false)
	})
}
