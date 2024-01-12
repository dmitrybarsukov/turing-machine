package validator

import (
	"testing"
)

func TestCountOfNumber(t *testing.T) {
	t.Run("has 3 of '3'", func(t *testing.T) {
		tt := newTester(t, CountOfNumber(3, []int{3})[0])

		tt.test(155, false)
		tt.test(322, false)
		tt.test(333, true)
		tt.test(422, false)
		tt.test(511, false)
	})

	t.Run("has 2 of '2'", func(t *testing.T) {
		tt := newTester(t, CountOfNumber(2, []int{2})[0])

		tt.test(155, false)
		tt.test(244, false)
		tt.test(333, false)
		tt.test(422, true)
		tt.test(224, true)
		tt.test(511, false)
	})

	t.Run("has 1 of '5'", func(t *testing.T) {
		tt := newTester(t, CountOfNumber(5, []int{1})[0])

		tt.test(155, false)
		tt.test(244, false)
		tt.test(353, true)
		tt.test(422, false)
		tt.test(511, true)
	})

	t.Run("has 0 of '3'", func(t *testing.T) {
		tt := newTester(t, CountOfNumber(3, []int{0})[0])

		tt.test(155, true)
		tt.test(244, true)
		tt.test(353, false)
		tt.test(422, true)
		tt.test(511, true)
	})
}

func TestCountOfParity(t *testing.T) {
	t.Run("has 3 of odd", func(t *testing.T) {
		tt := newTester(t, CountOfParity(Odd, []int{3})[0])

		tt.test(155, true)
		tt.test(322, false)
		tt.test(333, true)
		tt.test(433, false)
		tt.test(512, false)
	})

	t.Run("has 2 of even", func(t *testing.T) {
		tt := newTester(t, CountOfParity(Even, []int{2})[0])

		tt.test(155, false)
		tt.test(322, true)
		tt.test(333, false)
		tt.test(433, false)
		tt.test(542, true)
	})

	t.Run("has 1 of even", func(t *testing.T) {
		tt := newTester(t, CountOfParity(Even, []int{1})[0])

		tt.test(155, false)
		tt.test(322, false)
		tt.test(333, false)
		tt.test(433, true)
		tt.test(512, true)
	})
}

func TestParityComparator(t *testing.T) {
	t.Run("has more even than odd", func(t *testing.T) {
		tt := newTester(t, HasMoreNumbersWithParity([]Parity{Even})[0])

		tt.test(155, false)
		tt.test(322, true)
		tt.test(333, false)
		tt.test(433, false)
		tt.test(512, false)
	})

	t.Run("has more odd than even", func(t *testing.T) {
		tt := newTester(t, HasMoreNumbersWithParity([]Parity{Odd})[0])

		tt.test(521, true)
		tt.test(522, false)
		tt.test(123, true)
		tt.test(251, true)
		tt.test(241, false)
	})
}

func TestRepetitionCounter(t *testing.T) {
	t.Run("has no repetitions", func(t *testing.T) {
		tt := newTester(t, HasSomeRepeatingNumbers([]int{0})[0])

		tt.test(155, false)
		tt.test(322, false)
		tt.test(333, false)
		tt.test(433, false)
		tt.test(512, true)
	})

	t.Run("has 1 repetition", func(t *testing.T) {
		tt := newTester(t, HasSomeRepeatingNumbers([]int{1})[0])

		tt.test(155, true)
		tt.test(322, true)
		tt.test(333, false)
		tt.test(433, true)
		tt.test(512, false)
	})
}

func TestHasPairOfSameNumbers(t *testing.T) {
	t.Run("has no same number pairs", func(t *testing.T) {
		tt := newTester(t, PairOfNumbersExist([]bool{false})[0])

		tt.test(155, false)
		tt.test(322, false)
		tt.test(333, true)
		tt.test(433, false)
		tt.test(512, true)
	})

	t.Run("has some pairs", func(t *testing.T) {
		tt := newTester(t, PairOfNumbersExist([]bool{true})[0])

		tt.test(155, true)
		tt.test(232, true)
		tt.test(333, false)
		tt.test(112, true)
		tt.test(512, false)
	})
}
