package stalw

import (
	"math"
	"sort"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type Number interface {
	constraints.Float | constraints.Integer
}

// Min returns the minimum number of a series
func Min[T Number](nums ...T) T {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

// Max returns the maximum number of a series
func Max[T Number](nums ...T) T {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// Abs returns the absolute value of num
func Abs[T Number](num T) T {
	if num >= 0 {
		return num
	}
	return -num
}

// Constrain constraints num to a given range
func Constrain[T Number](num, min, max T) T {
	switch {
	case num < min:
		return min
	case num > max:
		return max
	default:
		return num
	}
}

// Map proportionally recalculates and returns a number from one range to another.
// If in one of the ranges its boundaries are inverted relative to another range, this will be taken into account.
// For example, Map(6, 0, 10, 0, 20) = 12, but Map(6, 0, 10, 20, 0) = 8
func Map[T Number](num, fromStart, fromEnd, toStart, toEnd T) float64 {
	relative := (float64(fromEnd) - float64(num)) / (float64(fromEnd) - float64(fromStart))
	newVal := float64(toEnd) - relative*(float64(toEnd)-float64(toStart))
	return newVal
}

// LinearXY returns the value of a function that is specified linearly by its argument.
// In other words, it returns the Y-coordinate of a point with an X-coordinate
// that lies on a line defined by a segment with points (x1, y1) and (x2, y2).
func LinearXY[T Number](x, x1, y1, x2, y2 T) float64 {
	return Map[T](x, x1, x2, y1, y2)
}

// Mean returns the average of a list of numbers.
// Note that if the numbers in the list are too large,
// the result may be incorrect due to the sum being overflowed.
func Mean[T Number](nums []T) float64 {
	var sum T
	for _, num := range nums {
		sum += num
	}

	res := float64(sum) / float64(len(nums))
	return res
}

// Median returns the median of a number series,
// using generics to process any float or int numbers
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

// Nearest returns a number equal to or closest to the original value and its index.
// In case of multiple results, the number with the lower index will be returned.
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
		// max(a,b) - min(a,b) work correct for large numbers, unlike abs(a-b)
		if d := Max[T](input[i], num) - Min[T](input[i], num); d < diff {
			diff = d
			index = i
		}
	}

	return input[index], index
}

// MostFrequent returns a list of the single most frequently occurring number,
// or a sorted list of such numbers if the number of occurrences matches.
// If input len equal zero, returns nil
func MostFrequent[T Number](nums []T) []T {
	ln := len(nums)
	if ln == 0 {
		return nil
	}

	maxCnt := 0
	stats := make(map[T]int)

	for _, num := range nums {
		stats[num]++
		if n := stats[num]; n > maxCnt {
			maxCnt = n
			if maxCnt > ln/2 {
				break
			}
		}
	}

	res := []T{}
	for num, cnt := range stats {
		if cnt == maxCnt {
			res = append(res, num)
		}
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

// Round rounds the number to the specified digits count after the period
func Round(num float64, digits int) float64 {
	precision := 1.0
	for digits > 0 {
		precision *= 10
		digits--
	}
	return math.Round(num*precision) / precision
}

// RoundMultiple rounds the number to a multiple of another number
func RoundMultiple[T Number](num T, multi T) T {
	if multi == 0 {
		return num
	}

	mf := float64(multi)
	quot := float64(num) / mf
	qRound := math.Round(quot)
	nRound := qRound * mf
	res := T(nRound)

	return res
}
