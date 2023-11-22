package math

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer
}

// Abs returns the absolute value of num
func Abs[T Number](num T) T {
	if num >= 0 {
		return num
	}
	return -num
}

// Min returns the minimum number of a series
func Min[T Number](nums ...T) T {
	var min *T
	for _, num := range nums {
		if min == nil || num < *min {
			min = &num
		}
	}
	return *min
}

// Max returns the maximum number of a series
func Max[T Number](nums ...T) T {
	var max *T
	for _, num := range nums {
		if max == nil || num > *max {
			max = &num
		}
	}
	return *max
}
