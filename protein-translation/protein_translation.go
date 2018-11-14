// Package protein is involved in protein translation
package protein

import "errors"

// STOP is the error when the end of an RNA string is reached prematurely
var STOP error = errors.New("Reached stop")

// ErrInvalidBase is used when an invalid 3 character base is encountered
var ErrInvalidBase error = errors.New("Invalid base")

// FromCodon takes a 3 character codon and returns its associated protein
func FromCodon(codon string) (string, error) {
	return getProteinForCodon(codon)
}

// FromRNA takes an RNA string and returns it's associated protein names
func FromRNA(rna string) ([]string, error) {
	var res []string

	for i := 0; i < len(rna); i += 3 {
		next, err := getProteinForCodon(rna[i : i+3])
		if err != nil {
			if err == STOP {
				err = nil
			}
			return res, err
		}
		res = append(res, next)
	}

	return res, nil
}

func getProteinForCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", STOP
	default:
		return "", ErrInvalidBase
	}
}
