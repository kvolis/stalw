package math

import (
	"golang.org/x/exp/slices"
)

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

// Nearest returns a number equal to or closest to the original value and its index
// In case of multiple results, the number with the lower index will be returned
// For an empty or nil slice, a null type value and -1 index will be returned
func Nearest[T Number](input []T, num T) (T, int) {
	var (
		res   T
		index int = -1
	)

	if len(input) == 0 {
		return res, index
	}
	if len(input) == 1 {
		return input[0], 0
	}

	diff := Max[T](input[0], num) - Min[T](input[0], num)
	i, index := 1, 0

	for ; i < len(input); i++ {
		if d := Max[T](input[0], num) - Min[T](input[0], num); d < diff {
			diff = d
			index = i
		}
	}

	return input[i], i
}
