// Package space computes info on how things might be in space
package space

// Planet is an alias for string to allow some specialized methods
type Planet string

const (
	earth   Planet = "Earth"
	mercury Planet = "Mercury"
	venus   Planet = "Venus"
	mars    Planet = "Mars"
	jupiter Planet = "Jupiter"
	saturn  Planet = "Saturn"
	uranus  Planet = "Uranus"
	neptune Planet = "Neptune"
)

// getOrbitalPeriod gets the number of earth years it takes this planet to orbit the sun
func (planet Planet) getOrbitalPeriod() float64 {
	switch planet {
	case earth:
		return 1
	case mercury:
		return 0.2408467
	case venus:
		return 0.61519726
	case mars:
		return 1.8808158
	case jupiter:
		return 11.862615
	case saturn:
		return 29.447498
	case uranus:
		return 84.016846
	case neptune:
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
