/**
* Space package computes info on how things might be in space
*/
package space

type Planet string

const (
	Earth Planet = "Earth"
	Mercury Planet = "Mercury"
	Venus Planet = "Venus"
	Mars Planet = "Mars"
	Jupiter Planet = "Jupiter"
	Saturn Planet = "Saturn"
	Uranus Planet = "Uranus"
	Neptune Planet = "Neptune"
)

// getOrbitalPeriod gets the number of earth years it takes this planet to orbit the sun
func (planet Planet) getOrbitalPeriod() float64 {
	switch planet {
	case Earth:
		return 1
	case Mercury:
		return 0.2408467
	case Venus:
		return 0.61519726
	case Mars:
		return 1.8808158
	case Jupiter:
		return 11.862615
	case Saturn:
		return 29.447498
	case Uranus:
		return 84.016846
	case Neptune:
		return 164.79132
	}

	panic("Could not find orbital period for " + planet)
}

const earthYearInSeconds = 365.25 * 60 * 60 * 24

// Age tells us how many times something has orbited the sun (its age in that planet's years)
// based on the number of seconds it has been in existence
func Age(seconds float64, planet Planet) float64 {
	return seconds / (planet.getOrbitalPeriod() * earthYearInSeconds)
}