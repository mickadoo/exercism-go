// Package hamming is used to calculate hamming distances
package hamming

import (
	"errors"
)

// Distance gets the hamming distance between two DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Strands are not equal length")
	}

	var distance int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
