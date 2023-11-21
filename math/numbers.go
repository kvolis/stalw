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
