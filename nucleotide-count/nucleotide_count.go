// Package dna handles dna related processesing
package dna

import (
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA []rune

const (
	adenine  = 'A'
	cytosine = 'C'
	guanine  = 'G'
	thymine  = 'T'
)

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{
		adenine:  0,
		cytosine: 0,
		guanine:  0,
		thymine:  0,
	}

	for _, nucleotide := range d {
		if _, ok := h[nucleotide]; !ok {
			return h, fmt.Errorf("invalid nucleotide: '%U'", nucleotide)
		}

		h[nucleotide]++
	}

	return h, nil
}
