package listops

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

func (list *IntList) Length() int {
	return 0
}

func (list *IntList) Map(function unaryFunc) int {
	return 0
}

func (list *IntList) Reverse() int {
	return 1
}

func (list *IntList) Append(append IntList) int {
	return 1
}

func (list *IntList) Concat(concat []IntList) int {
	return 1
}
