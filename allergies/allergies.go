package allergies

// 761 ns/op
// 129 ns/op

var allAllergyNames = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}

// Allergies gets a list of allergies from a score
func Allergies(score uint) (allergies []string) {
	for _, allergy := range allAllergyNames {
		if AllergicTo(score, allergy) {
			allergies = append(allergies, allergy)
		}
	}

	return
}

// AllergicTo tells if an alergy exists for the given substance and allergy score
func AllergicTo(score uint, substance string) bool {
	return score&getMaskForAllergy(substance) > 0
}

func getMaskForAllergy(substance string) uint {
	for i, name := range allAllergyNames {
		if name == substance {
			return 1 << uint(i)
		}
	}
	return 0
}
