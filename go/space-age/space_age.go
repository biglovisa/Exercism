package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	earth := float64(31557600)
	divisor := map[Planet]float64{
		"Mercury": earth * 0.2408467,
		"Venus": earth * 0.61519726,
		"Earth": earth,
		"Mars": earth * 1.8808158,
		"Jupiter": earth * 11.862615,
		"Saturn": earth * 29.447498,
		"Uranus": earth * 84.016846,
		"Neptune": earth * 164.79132,
	}[planet]

	return seconds / divisor
}