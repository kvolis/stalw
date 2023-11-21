package math

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type Number interface {
	constraints.Float | constraints.Integer
}

// Median returns the median of a number series, using generics to process any float or int numbers
func Median[T Number](input []T) float64 {
	ln := len(input)

	// input will not be changed
	sub := make([]T, ln)
	copy(sub, input)
	slices.Sort(sub)

	res := float64(sub[ln/2])

	if ln%2 == 0 {
		// a/2 + b/2 work correct for large numbers, unlike (a+b)/2
		res = res/2 + float64(sub[ln/2-1])/2
	}

	return res
}
