package grains

import (
	"errors"
	"math"
)

// Square returns the number of grains on a certain square
func Square(square int) (count uint64, err error) {
	if square > 64 || square < 1 {
		return count, errors.New("square must be between 1 and 64")
	}

	return uint64(math.Pow(2, float64(square-1))), nil
}

// Total returns the total of all grains of rice on the board
func Total() uint64 {
	var expected uint64
	for i := 1; i < 65; i++ {
		current, _ := Square(i)
		expected += current
	}

	return expected
}
