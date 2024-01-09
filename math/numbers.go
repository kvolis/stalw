package math

import (
	"golang.org/x/exp/constraints"
)

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
