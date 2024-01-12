package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCombinations(t *testing.T) {
	result := Combinations([]int{1, 2}, 3)
	expected := [][]int{
		{1, 1, 1},
		{1, 1, 2},
		{1, 2, 1},
		{1, 2, 2},
		{2, 1, 1},
		{2, 1, 2},
		{2, 2, 1},
		{2, 2, 2},
	}
	require.Equal(t, expected, result)
}
