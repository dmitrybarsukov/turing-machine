package validator

import (
	"testing"
)

func TestCountChecker(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		t.Run("has 3 of '3'", func(t *testing.T) {
			tt := newTester(t, countChecker{Checker: equalityChecker{Number: 3}, Count: 3})
			tt.test(155, false)
			tt.test(322, false)
			tt.test(333, true)
			tt.test(422, false)
			tt.test(511, false)
		})

		t.Run("has 2 of '2'", func(t *testing.T) {
			tt := newTester(t, countChecker{Checker: equalityChecker{Number: 2}, Count: 2})

			tt.test(155, false)
			tt.test(244, false)
			tt.test(333, false)
			tt.test(422, true)
			tt.test(224, true)
			tt.test(511, false)
		})

		t.Run("has 1 of '5'", func(t *testing.T) {
			tt := newTester(t, countChecker{Checker: equalityChecker{Number: 5}, Count: 1})

			tt.test(155, false)
			tt.test(244, false)
			tt.test(353, true)
			tt.test(422, false)
			tt.test(511, true)
		})

		t.Run("has 0 of '3'", func(t *testing.T) {
			tt := newTester(t, countChecker{Checker: equalityChecker{Number: 3}, Count: 0})

			tt.test(155, true)
			tt.test(244, true)
			tt.test(353, false)
			tt.test(422, true)
			tt.test(511, true)
		})
	})

	t.Run("parity", func(t *testing.T) {
		t.Run("has 3 of odd", func(t *testing.T) {
			tt := newTester(t, countChecker{Checker: parityChecker{Parity: Odd}, Count: 3})

			tt.test(155, true)
			tt.test(322, false)
			tt.test(333, true)
			tt.test(433, false)
			tt.test(512, false)
		})

		t.Run("has 2 of even", func(t *testing.T) {
			tt := newTester(t, countChecker{Checker: parityChecker{Parity: Even}, Count: 2})

			tt.test(155, false)
			tt.test(322, true)
			tt.test(333, false)
			tt.test(433, false)
			tt.test(542, true)
		})

		t.Run("has 1 of even", func(t *testing.T) {
			tt := newTester(t, countChecker{Checker: parityChecker{Parity: Even}, Count: 1})

			tt.test(155, false)
			tt.test(322, false)
			tt.test(333, false)
			tt.test(433, true)
			tt.test(512, true)
		})
	})
}

func TestRepetitionCounter(t *testing.T) {
	t.Run("has no repetitions", func(t *testing.T) {
		tt := newTester(t, repetitionCounter{Result: 0})

		tt.test(155, false)
		tt.test(322, false)
		tt.test(333, false)
		tt.test(433, false)
		tt.test(512, true)
	})

	t.Run("has 1 repetition", func(t *testing.T) {
		tt := newTester(t, repetitionCounter{Result: 1})

		tt.test(155, true)
		tt.test(322, true)
		tt.test(333, false)
		tt.test(433, true)
		tt.test(512, false)
	})

	t.Run("has 2 repetition", func(t *testing.T) {
		tt := newTester(t, repetitionCounter{Result: 2})

		tt.test(155, false)
		tt.test(322, false)
		tt.test(333, true)
		tt.test(433, false)
		tt.test(512, false)
	})
}
