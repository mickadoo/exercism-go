// Package scrabble is used for calculating scrabble scores
package scrabble

import (
	"strings"
)

// Score computes the scrabble score for a word
func Score(word string) int {
	score := 0
	word = strings.ToUpper(word)

	for _, letter := range word {
		score += getLetterScore(letter)
	}

	return score
}

func getLetterScore(letter rune) int {
	switch letter {
	case 'A':
		return 1
	case 'B':
		return 3
	case 'C':
		return 3
	case 'D':
		return 2
	case 'E':
		return 1
	case 'F':
		return 4
	case 'G':
		return 2
	case 'H':
		return 4
	case 'I':
		return 1
	case 'J':
		return 8
	case 'K':
		return 5
	case 'L':
		return 1
	case 'M':
		return 3
	case 'N':
		return 1
	case 'O':
		return 1
	case 'P':
		return 3
	case 'Q':
		return 10
	case 'R':
		return 1
	case 'S':
		return 1
	case 'T':
		return 1
	case 'U':
		return 1
	case 'V':
		return 4
	case 'W':
		return 4
	case 'X':
		return 8
	case 'Y':
		return 4
	case 'Z':
		return 10
	}

	panic("Unknown letter")
}
