package internal

import "math"

type Coord struct {
	X float64
	Y float64
}

// sqrt((x2-x1)^2 + (y2-y1)^2)
func CalculateCost(start, end *Coord) float64 {
	return math.Sqrt(math.Pow(end.X-start.X, 2) + math.Pow(end.Y-start.Y, 2))
}
