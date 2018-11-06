package listops

// IntList is a list of integers
type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Foldr for args [x1, x2, ..., xn] returns f(x1, f(x2, ..., f(xn, init)...))
func (l IntList) Foldr(function binFunc, initial int) int {
	current := initial
	for i := len(l) - 1; i >= 0; i-- {
		current = function(l[i], current)
	}

	return current
}

// Foldl for args [x1, x2, ..., xn] returns f(xn,...,f(x2, f(x1, init))...)
func (l IntList) Foldl(function binFunc, initial int) int {
	current := initial
	for _, val := range l {
		current = function(current, val)
	}

	return current
}

// Filter goes over each element, checks if it passed a check function and returns those elements that do
func (l IntList) Filter(function predFunc) IntList {
	result := IntList{}

	for _, val := range l {
		if function(val) {
			result = append(result, val)
		}
	}

	return result
}

// Length gets the length of the list. Not sure if I'm allowed use len() here
func (l IntList) Length() int {
	len := 0
	for range l {
		len++
	}

	return len
}

// Map applies a function on all members of a list
func (l IntList) Map(function unaryFunc) IntList {
	for index, val := range l {
		l[index] = function(val)
	}

	return l
}

// Reverse reverses a list
func (l IntList) Reverse() IntList {
	for i := 0; i < len(l)/2; i++ {
		tmp := l[i]
		oppositeIndex := len(l) - i - 1
		l[i] = l[oppositeIndex]
		l[oppositeIndex] = tmp
	}

	return l
}

// Append add an element to a list.
func (l IntList) Append(toAdd IntList) IntList {
	combined := make(IntList, len(l)+len(toAdd))
	i := 0
	for _, val := range l {
		combined[i] = val
		i++
	}
	for _, val := range toAdd {
		combined[i] = val
		i++
	}

	return combined
}

// Concat takes a slice of integer lists and adds all elements from them to this list
func (l IntList) Concat(concat []IntList) IntList {
	for _, nextList := range concat {
		l = l.Append(nextList)
	}

	return l
}
