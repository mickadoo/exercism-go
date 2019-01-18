package grains

import (
	"errors"
)

// Square returns the number of grains on a certain square
func Square(square int) (count uint64, err error) {
	if square > 64 || square < 1 {
		return count, errors.New("square must be between 1 and 64")
	}

	return 1 << uint(square-1), nil
}

// Total returns the total of all grains of rice on the board
func Total() uint64 {
	return 1<<64 - 1
}
