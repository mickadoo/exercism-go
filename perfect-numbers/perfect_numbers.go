package perfect

import "errors"

// Classification is a type of number
type Classification int

// ErrOnlyPositive is returned when you attempt to classify a non positive number
var ErrOnlyPositive = errors.New("Cannot classify a non-positive number")

const (
	// ClassificationDeficient is when the sum of the factors are less than the number
	ClassificationDeficient = iota
	// ClassificationPerfect is when the sum of the factors is equal to the number
	ClassificationPerfect
	// ClassificationAbundant is when the sum of the factors is greater than the number
	ClassificationAbundant
)

// Classify classifies the input into perfect, abuntant or deficient
func Classify(input int64) (class Classification, err error) {
	if input < 1 {
		return class, ErrOnlyPositive
	} else if input == 1 {
		return ClassificationDeficient, nil
	}

	total := getSumOfFactors(input)

	switch {
	case total == input:
		class = ClassificationPerfect
	case total > input:
		class = ClassificationAbundant
	case total < input:
		class = ClassificationDeficient
	}

	return
}

func getSumOfFactors(num int64) int64 {
	total := int64(1)
	// we're only looping up to the square root of num
	for i := int64(2); i*i <= num; i++ {
		if num%i == 0 {
			total += i
			// since i is a factor, num/i is also a factor
			// unless they're the same, e.g. 2*2 for 4)
			if num/i != i {
				total += num / i
			}
		}
	}

	return total
}
