// Package strain helps filter lists
package strain

// Ints is a list of integers
type Ints []int

// Lists is a list of integer lists
type Lists [][]int

// Strings is a list of strings
type Strings []string

// Keep returns a list of all members of the list that satisfy the predicate
func (ints Ints) Keep(predicate func(int) bool) Ints {
	var result Ints
	for _, val := range ints {
		if predicate(val) {
			result = append(result, val)
		}
	}

	return result
}

// Discard returns a list of all members of the list that do not satisfy the predicate
func (ints Ints) Discard(predicate func(int) bool) Ints {
	return ints.Keep(func(i int) bool {
		return !predicate(i)
	})
}

// Keep returns a list of all members of the list that satisfy the predicate
func (lists Lists) Keep(predicate func([]int) bool) Lists {
	var result Lists
	for _, val := range lists {
		if predicate(val) {
			result = append(result, val)
		}
	}

	return result
}

// Keep returns a list of all members of the list that satisfy the predicate
func (strs Strings) Keep(predicate func(string) bool) Strings {
	var result Strings
	for _, val := range strs {
		if predicate(val) {
			result = append(result, val)
		}
	}

	return result
}
