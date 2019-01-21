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
	}

	factors := findFactors(input)

	var total int64
	for _, factor := range factors {
		total += factor
	}

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

func findFactors(num int64) (factors []int64) {
	for i := int64(1); i <= num/2; i++ {
		if num%i == 0 {
			factors = append(factors, i)
		}
	}

	return
}
