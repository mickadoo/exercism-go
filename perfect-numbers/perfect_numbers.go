package perfect

import (
	"errors"
	"math"
)

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

	return getSumOfFactors(input), nil
}

func getSumOfFactors(num int64) Classification {
	total := int64(1)
	sqrRoot := int64(math.Sqrt(float64(num)))
	var otherFactor int64

	// we only need to check as far as square root to get unique factors using this method
	for factor := int64(2); factor <= sqrRoot; factor++ {
		if num%factor == 0 {
			total += factor
			// since i is a factor, num/i is also a factor
			// unless they're the same, e.g. 2*2 for 4)
			otherFactor = num / factor
			if otherFactor != factor {
				total += otherFactor
			}
		}

		if total > num {
			// shortcut, if total already exceeds num we know the type
			return ClassificationAbundant
		}
	}

	switch {
	case total == num:
		return ClassificationPerfect
	case total > num:
		return ClassificationAbundant
	case total < num:
		return ClassificationDeficient
	default:
		panic("Should not be able to reach this point")
	}
}
