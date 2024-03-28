package stalw

// Count returns a count of occurrences of an element in a list of elements
func Count[T comparable](elements []T, element T) int {
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

// FitTo returns the result outputW and outputH of scaling the
// specified inputW and inputH into the target targetW and targetH.
// In other words, this is the result of fitting in
// such a way as not to go beyond the target boundaries.
func FitTo[T Number](inputW, inputH, targetW, targetH T) (outputW, outputH T) {
	wScale := float64(targetW) / float64(inputW)
	hScale := float64(targetH) / float64(inputH)

	scale := Min(wScale, hScale)

	outputW = T(float64(inputW) * scale)
	outputH = T(float64(inputH) * scale)
	return
}
