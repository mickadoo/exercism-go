// Package strand deals with DNA strand processing
package strand

type strand []rune

const (
	adenine  rune = 'A'
	cytosine      = 'C'
	guanine       = 'G'
	thymine       = 'T'
	uracil        = 'U'
)

var dnaToRnaMap = map[rune]rune{
	guanine:  cytosine,
	cytosine: guanine,
	thymine:  adenine,
	adenine:  uracil,
}

// ToRNA returns the RNA strand for a given DNA strand
func ToRNA(dna string) string {
	var rna strand
	for i := 0; i < len(dna); i++ {
		rna = append(rna, dnaToRnaMap[rune(dna[i])])
	}
	return string(rna)
}
