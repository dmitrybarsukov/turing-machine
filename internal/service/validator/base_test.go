package validator

import (
	"testing"

	"github.com/dmitrybarsukov/turing-machine/internal/domain"

	"github.com/stretchr/testify/require"
)

func codeFromInt(number int) domain.Code {
	return domain.Code{number / 100 % 10, number / 10 % 10, number % 10}
}

type validatorTester struct {
	t         *testing.T
	validator domain.Validator
}

func newTester(t *testing.T, validator domain.Validator) validatorTester {
	return validatorTester{t: t, validator: validator}
}

func (t validatorTester) test(code int, expected bool) {
	result := t.validator.Validate(codeFromInt(code))
	require.Equalf(t.t, expected, result, "validator {%v} on code {%v}", t.validator, code)
}
