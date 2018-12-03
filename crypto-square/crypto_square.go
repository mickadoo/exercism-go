package cryptosquare

import (
	"math"
	"unicode"
)

// Encode returns the "square code" encoded version of the input
func Encode(input string) string {
	input = normalize(input)
	colCount, rowCount := getRowAndColumnCount(len(input))
	result := ""

	for col := 0; col < colCount; col++ {
		for row := 0; row < rowCount; row++ {
			// get the next character from the column
			index := col + (row * colCount)

			// with uneven row and columns this can be empty
			if index < len(input) {
				result += string(input[index])
			} else {
				result += " "
			}
		}
		if col+1 < colCount {
			// space between column, except on last
			result += " "
		}
	}

	return result
}

func getRowAndColumnCount(length int) (columns int, rows int) {
	rows = int(math.Round(math.Sqrt(float64(length))))
	if rows*rows >= length {
		columns = rows
	} else {
		columns = rows + 1
	}

	return
}

func normalize(input string) string {
	var result string
	for _, char := range input {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			result += string(unicode.ToLower(char))
		}
	}

	return result
}
