//Package etl is for extracting, transforming and loading - converting legacy data to a new format
package etl

import (
	"strings"
)

// Transform takes old scoring data and transforms it to the new format
func Transform(oldSystem map[int][]string) map[string]int {
	newSystem := make(map[string]int)

	for score, letters := range oldSystem {
		for _, letter := range letters {
			newSystem[strings.ToLower(letter)] = score
		}
	}

	return newSystem
}
