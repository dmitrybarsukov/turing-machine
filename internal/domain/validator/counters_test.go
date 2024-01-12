package validator

import (
	"testing"
)

func TestCountOfNumber(t *testing.T) {
	t.Run("has 3 of '3'", func(t *testing.T) {
		tt := newTester(t, countOfNumberChecker{Number: 3, count: 3})
		tt.test(155, false)
		tt.test(322, false)
		tt.test(333, true)
		tt.test(422, false)
		tt.test(511, false)
	})

	t.Run("has 2 of '2'", func(t *testing.T) {
		tt := newTester(t, countOfNumberChecker{Number: 2, count: 2})

		tt.test(155, false)
		tt.test(244, false)
		tt.test(333, false)
		tt.test(422, true)
		tt.test(224, true)
		tt.test(511, false)
	})

	t.Run("has 1 of '5'", func(t *testing.T) {
		tt := newTester(t, countOfNumberChecker{Number: 5, count: 1})

		tt.test(155, false)
		tt.test(244, false)
		tt.test(353, true)
		tt.test(422, false)
		tt.test(511, true)
	})

	t.Run("has 0 of '3'", func(t *testing.T) {
		tt := newTester(t, countOfNumberChecker{Number: 3, count: 0})

		tt.test(155, true)
		tt.test(244, true)
		tt.test(353, false)
		tt.test(422, true)
		tt.test(511, true)
	})
}

func TestCountOfParity(t *testing.T) {
	t.Run("has 3 of odd", func(t *testing.T) {
		tt := newTester(t, countOfParityChecker{Parity: Odd, count: 3})

		tt.test(155, true)
		tt.test(322, false)
		tt.test(333, true)
		tt.test(433, false)
		tt.test(512, false)
	})

	t.Run("has 2 of even", func(t *testing.T) {
		tt := newTester(t, countOfParityChecker{Parity: Even, count: 2})

		tt.test(155, false)
		tt.test(322, true)
		tt.test(333, false)
		tt.test(433, false)
		tt.test(542, true)
	})

	t.Run("has 1 of even", func(t *testing.T) {
		tt := newTester(t, countOfParityChecker{Parity: Even, count: 1})

		tt.test(155, false)
		tt.test(322, false)
		tt.test(333, false)
		tt.test(433, true)
		tt.test(512, true)
	})
}

func TestParityComparator(t *testing.T) {
	t.Run("has more even than odd", func(t *testing.T) {
		tt := newTester(t, parityCountComparator{result: Even})

		tt.test(155, false)
		tt.test(322, true)
		tt.test(333, false)
		tt.test(433, false)
		tt.test(512, false)
	})

	t.Run("has more odd than even", func(t *testing.T) {
		tt := newTester(t, parityCountComparator{result: Odd})

		tt.test(521, true)
		tt.test(522, false)
		tt.test(123, true)
		tt.test(251, true)
		tt.test(241, false)
	})
}

func TestRepetitionCounter(t *testing.T) {
	t.Run("has no repetitions", func(t *testing.T) {
		tt := newTester(t, repetitionCounter{result: 0})

		tt.test(155, false)
		tt.test(322, false)
		tt.test(333, false)
		tt.test(433, false)
		tt.test(512, true)
	})

	t.Run("has 1 repetition", func(t *testing.T) {
		tt := newTester(t, repetitionCounter{result: 1})

		tt.test(155, true)
		tt.test(322, true)
		tt.test(333, false)
		tt.test(433, true)
		tt.test(512, false)
	})

	t.Run("has 2 repetition", func(t *testing.T) {
		tt := newTester(t, repetitionCounter{result: 2})

		tt.test(155, false)
		tt.test(322, false)
		tt.test(333, true)
		tt.test(433, false)
		tt.test(512, false)
	})
}

func TestHasPairOfSameNumbers(t *testing.T) {
	t.Run("has no same number pairs", func(t *testing.T) {
		tt := newTester(t, hasPairOfSameNumbersChecker{result: false})

		tt.test(155, false)
		tt.test(322, false)
		tt.test(333, true)
		tt.test(433, false)
		tt.test(512, true)
	})

	t.Run("has some pairs", func(t *testing.T) {
		tt := newTester(t, hasPairOfSameNumbersChecker{result: true})

		tt.test(155, true)
		tt.test(232, true)
		tt.test(333, false)
		tt.test(112, true)
		tt.test(512, false)
	})
}
