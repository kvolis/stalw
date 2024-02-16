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
