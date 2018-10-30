// Package scale is used for generating musical scales
package scale

import (
	"fmt"
)

type pitch struct {
	sharpNotation string
	flatNotation  string
	next          *pitch
}

func (p pitch) toString(notationType string) string {
	switch notationType {
	case "N":
		return p.sharpNotation
	case "S":
		return p.sharpNotation
	case "F":
		return p.flatNotation
	}

	panic("Could not handle notation type" + notationType)
}

type chromaticScale struct {
	current *pitch
}

func (s *chromaticScale) getNext() *pitch {
	s.current = s.current.next

	return s.current
}

func (s *chromaticScale) moveTo(note string) {
	for {
		if s.current.flatNotation == note || s.current.sharpNotation == note {
			return
		}
		s.getNext()
	}
}

func (s *chromaticScale) getSequence(tonic string, length int) []string {
	result := make([]string, length)
	s.moveTo(tonic)
	notationType := getPitchNotationType(tonic)

	fmt.Println("Notation type: " + notationType)

	for i := 0; i < length; i++ {
		result[i] = s.current.toString(notationType)
		s.getNext()
	}

	return result
}

// Scale returns a musical scale from the provided tonic and interval
func Scale(tonic string, interval string) []string {
	scale := createScale()

	return scale.getSequence(tonic, 12)
}

func getPitchNotationType(tonic string) string {
	for _, match := range []string{"C", "a"} {
		if tonic == match {
			// none
			return "N"
		}
	}

	for _, match := range []string{"G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#"} {
		if tonic == match {
			// sharps
			return "S"
		}
	}

	// flats
	return "F"
}

func createScale() *chromaticScale {
	ab := pitch{"G#", "Ab", nil}
	g := pitch{"G", "G", &ab}
	gb := pitch{"F#", "Gb", &g}
	f := pitch{"F", "F", &gb}
	e := pitch{"E", "E", &f}
	eb := pitch{"D#", "Eb", &e}
	d := pitch{"D", "D", &eb}
	db := pitch{"C#", "Db", &d}
	c := pitch{"C", "C", &db}
	b := pitch{"B", "B", &c}
	bb := pitch{"A#", "Bb", &b}
	a := pitch{"A", "A", &bb}
	ab.next = &a

	return &chromaticScale{&a}
}
