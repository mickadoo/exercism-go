// Package accumulate is for doing accumulative operations on lists
package accumulate

// Making this a type just to make it more readable
type converterFunc func(string) string

// Accumulate runs a given function on a list
func Accumulate(given []string, converter converterFunc) []string {
	for index, value := range given {
		given[index] = converter(value)
	}

	return given
}
