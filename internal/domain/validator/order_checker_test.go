package validator

import "testing"

func TestOrderChecker(t *testing.T) {
	t.Run("Asc", func(t *testing.T) {
		tt := newTester(t, orderChecker{result: Ascending})

		tt.test(435, false)
		tt.test(322, false)
		tt.test(333, false)
		tt.test(432, false)
		tt.test(345, true)
		tt.test(355, false)
	})

	t.Run("Desc", func(t *testing.T) {
		tt := newTester(t, orderChecker{result: Descending})

		tt.test(435, false)
		tt.test(321, true)
		tt.test(333, false)
		tt.test(432, true)
		tt.test(345, false)
		tt.test(522, false)
	})

	t.Run("Chaotic", func(t *testing.T) {
		tt := newTester(t, orderChecker{result: Unordered})

		tt.test(435, true)
		tt.test(321, false)
		tt.test(333, true)
		tt.test(432, false)
		tt.test(345, false)
		tt.test(522, true)
	})
}
