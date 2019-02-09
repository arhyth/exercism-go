/*
package space exposes the Age function
coz you know people always want to know
how old you are
*/
package space

type Planet string

const (
	Earth Planet = "Earth"
	Jupiter Planet = "Jupiter"
	Mars Planet = "Mars"
	Mercury Planet = "Mercury"
	Neptune Planet = "Neptune"
	Saturn Planet = "Saturn"
	Uranus Planet = "Uranus"
	Venus Planet = "Venus"
)

var earthMultiple = map[Planet]float64{
	Jupiter: 11.862615,
	Mars: 1.8808158,
	Mercury: 0.2408467,
	Neptune: 164.79132,
	Saturn: 29.447498,
	Uranus: 84.016846,
	Venus: 0.61519726,
}

const earthYearSeconds int64 = 31557600

// Age computes a person's age in another planet
// why though, i don't know
func Age(seconds float64, planet Planet) float64 {
	ageEarth := seconds / float64(earthYearSeconds) 
	var age float64
	if planet == Earth {
		age = ageEarth
	} else {
		age = ageEarth / earthMultiple[planet]
	}
	return age
}

