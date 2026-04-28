package array

import "slices"

func Contains[T comparable](input []T, item T) bool {
	return slices.Contains(input, item)
}

func Map[T any, R any](input []T, fn func(T) R) []R {
	result := make([]R, len(input))
	for i, v := range input {
		result[i] = fn(v)
	}

	return result
}

func Filter[T any](input []T, fn func(T) bool) []T {
	var result []T
	for _, v := range input {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

func Every[T any](input []T, fn func(T) bool) bool {
	return !slices.ContainsFunc(input, func(v T) bool {
		return !fn(v)
	})
}

func Some[T any](input []T, fn func(T) bool) bool {
	return slices.ContainsFunc(input, fn)
}

func Find[T any](input []T, fn func(T) bool) *T {
	index := slices.IndexFunc(input, fn)
	if index == -1 {
		return nil
	}

	return &input[index]
}

func FindIndex[T any](input []T, fn func(T) bool) int {
	return slices.IndexFunc(input, fn)
}
