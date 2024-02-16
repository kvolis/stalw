package stalw

// Count returns a count of occurrences of an element in a list of elements
func Count[T comparable](element T, elements []T) int {
	cnt := 0
	for _, e := range elements {
		if e == element {
			cnt++
		}
	}
	return cnt
}

// MostFrequent returns a list of the single most frequently occurring element,
// or a list of such elements if the number of occurrences matches.
// If input len equal zero, returns nil
func MostFrequent[T comparable](elements []T) []T {
	ln := len(elements)
	if ln == 0 {
		return nil
	}

	maxCnt := 0
	stats := make(map[T]int)

	for _, num := range elements {
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
	return res
}
