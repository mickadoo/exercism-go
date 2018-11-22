// Package pythagorean is used for calculating pythagorean triplets
package pythagorean

// Triplet is an array of 3 unique numbers
type Triplet [3]int

func (t Triplet) sum() int {
	return t[0] + t[1] + t[2]
}

// Range calculates all pythagorean triples from the given min and max
func Range(min, max int) []Triplet {
	return getPythagoranTriplets(min, max)
}

// Sum finds all pythagorean triplets whose sum is equal to p
func Sum(p int) (result []Triplet) {
	triplets := getPythagoranTriplets(1, p-1)
	for _, triplet := range triplets {
		if triplet.sum() == p {
			result = append(result, triplet)
		}
	}

	return
}

func getPythagoranTriplets(min, max int) (result []Triplet) {
	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			for c := b + 1; c <= max; c++ {
				if (a*a)+(b*b) == (c * c) {
					result = append(result, Triplet{a, b, c})
				}
			}
		}
	}
	return
}
