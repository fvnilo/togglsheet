package export

import "math"

// CalculateTimeEntry convert milliseconds to hours and rounds it to the nearest quater.
func CalculateTimeEntry(ms int) float64 {
	hours := float64(ms) / 3600000
	roundedHours := math.Round(hours*4) / 4

	if roundedHours == 0 {
		roundedHours = 0.25
	}

	return roundedHours
}
