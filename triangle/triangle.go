// Package triangle is used in calculations about triangles
package triangle

import (
	"math"
)

// Kind is the type of available triangles
type Kind int

const (
	// NaT = "not a triangle", e.g. if one side is zero length
	NaT = iota
	// Equ = equilateral triangle, where all sides are the same length
	Equ
	// Iso = isosceles triangle, with at least two sides the same length
	Iso
	// Sca = scalene triangle, with all sides of different lengths
	Sca
)

// KindFromSides determines a triangle type based on the length of its 3 sides
func KindFromSides(a, b, c float64) Kind {
	if !isValidTriangle(a, b, c) {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}

func isValidTriangle(a, b, c float64) bool {
	if !isValidSide(a, b, c) {
		return false
	}

	// check triangle inequality
	if a+b < c || b+c < a || a+c < b {
		return false
	}

	return true
}

func isValidSide(sides ...float64) bool {
	for _, side := range sides {
		if side <= 0 || math.IsNaN(side) || math.IsInf(side, 1) || math.IsInf(side, -1) {
			return false
		}
	}
	return true
}
