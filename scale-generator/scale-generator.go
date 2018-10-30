// Package scale is used for generating musical scales
package scale

import (
	"strings"
)

const (
	chromaticScaleLenght = 12
	pitchNotationFlats   = 'F'
	pitchNotationSharps  = 'S'
	pitchNotationNone    = 'N'
)

type pitch struct {
	sharpNotation string
	flatNotation  string
	next          *pitch
}

func (p *pitch) toString(tonic string) string {
	notationType := getPitchNotationType(tonic)

	switch notationType {
	case pitchNotationNone:
		return p.sharpNotation
	case pitchNotationSharps:
		return p.sharpNotation
	case pitchNotationFlats:
		return p.flatNotation
	}

	panic("Could not handle notation type")
}

func (p *pitch) isEqual(note string) bool {
	return strings.ToLower(note) == strings.ToLower(p.sharpNotation) ||
		strings.ToLower(note) == strings.ToLower(p.flatNotation)
}

type chromaticScale struct {
	current *pitch
}

func (s *chromaticScale) next() {
	s.current = s.current.next
}

func (s *chromaticScale) skip(num int) {
	for num > 0 {
		s.next()
		num--
	}
}

func (s *chromaticScale) moveTo(note string) {
	for i := 0; i < chromaticScaleLenght; i++ {
		if s.current.isEqual(note) {
			return
		}
		s.next()
	}

	panic("Could not find note " + note + " in chromatic scale")
}

func (s *chromaticScale) getScale(tonic string, interval string) []string {
	result := make([]string, len(interval))
	s.moveTo(tonic)

	for i := 0; i < len(interval); i++ {
		result[i] = s.current.toString(tonic)

		switch interval[i] {
		case 'M':
			s.skip(2)
		case 'A':
			s.skip(3)
		case 'm':
			s.next()
		}
	}

	return result
}

// Scale returns a musical scale from the provided tonic and interval
func Scale(tonic string, interval string) []string {
	scale := createChromaticScale()

	if interval == "" {
		// create basic interval with no skips for chromatic scale
		interval = strings.Repeat("m", chromaticScaleLenght)
	}

	return scale.getScale(tonic, interval)
}

func getPitchNotationType(tonic string) byte {
	for _, match := range []string{"C", "a"} {
		if tonic == match {
			return pitchNotationNone
		}
	}

	for _, match := range []string{"G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#"} {
		if tonic == match {
			return pitchNotationSharps
		}
	}

	return pitchNotationFlats
}

func createChromaticScale() *chromaticScale {
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
