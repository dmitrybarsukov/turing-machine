package util

import "github.com/samber/lo"

func Cross[T any](variants [][]T, items []T) [][]T {
	if len(variants) == 0 {
		return lo.Map(items, func(it T, _ int) []T {
			return []T{it}
		})
	}

	return lo.FlatMap(variants, func(it []T, _ int) [][]T {
		result := make([][]T, 0, len(items))
		for _, i := range items {
			newSlice := make([]T, len(it), len(it)+1)
			copy(newSlice, it)
			newSlice = append(newSlice, i)
			result = append(result, newSlice)
		}
		return result
	})
}

func Combinations[T any](items []T, dimensions int) [][]T {
	var result [][]T
	for i := 0; i < dimensions; i++ {
		result = Cross(result, items)
	}

	return result
}

func Combine1[TIn, TOut any](vals []TIn, fn func(TIn) TOut) []TOut {
	return lo.Map(vals, func(it TIn, _ int) TOut {
		return fn(it)
	})
}

func Combine2[TIn1, TIn2, TOut any](vals1 []TIn1, vals2 []TIn2, fn func(TIn1, TIn2) TOut) []TOut {
	return lo.FlatMap(vals1, func(it1 TIn1, _ int) []TOut {
		return lo.Map(vals2, func(it2 TIn2, _ int) TOut {
			return fn(it1, it2)
		})
	})
}
