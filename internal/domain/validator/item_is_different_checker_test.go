package validator

import (
	"testing"
	"turing-machine/internal/domain"
)

func TestOneItemIsDifferent(t *testing.T) {
	t.Run("[0] is less", func(t *testing.T) {
		tt := newTester(t, itemIsDifferentChecker{Result: Less, item: domain.CodeItem0})

		tt.test(155, true)
		tt.test(322, false)
		tt.test(333, false)
		tt.test(433, false)
		tt.test(455, true)
	})

	t.Run("[1] is less", func(t *testing.T) {
		tt := newTester(t, itemIsDifferentChecker{Result: Less, item: domain.CodeItem1})

		tt.test(435, true)
		tt.test(322, false)
		tt.test(333, false)
		tt.test(432, false)
		tt.test(415, true)
	})

	t.Run("[2] is more", func(t *testing.T) {
		tt := newTester(t, itemIsDifferentChecker{Result: More, item: domain.CodeItem2})

		tt.test(435, true)
		tt.test(322, false)
		tt.test(333, false)
		tt.test(432, false)
		tt.test(415, true)
	})
}
