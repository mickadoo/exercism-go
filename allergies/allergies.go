package allergies

// 528 ns/op
// 73.6 ns/op

const (
	eggs = 1 << iota
	peanuts
	shellfish
	strawberries
	tomatoes
	chocolate
	pollen
	cats
)

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
	switch substance {
	case "eggs":
		return eggs
	case "peanuts":
		return peanuts
	case "shellfish":
		return shellfish
	case "strawberries":
		return strawberries
	case "tomatoes":
		return tomatoes
	case "chocolate":
		return chocolate
	case "pollen":
		return pollen
	case "cats":
		return cats
	default:
		return 0
	}
}
