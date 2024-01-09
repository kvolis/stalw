package math

import (
	"sort"

	"golang.org/x/exp/slices"
)

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
// If input len equal zero, returns nil.
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

// Count returns a count of occurrences of a number in a list of numbers
func Count[T Number](num T, nums []T) int {
	cnt := 0
	for _, n := range nums {
		if n == num {
			cnt++
		}
	}
	return cnt
}
