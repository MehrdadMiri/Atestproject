package command

import "math"

const (
	NextCycle      = iota
	NewDestination = iota
)

type Location struct {
	X, Y float64
}

type Command struct {
	Type        int
	Destination Location
}

type Distance float64

func (a Location) Distance(b Location) Distance {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return Distance(math.Sqrt((dx * dx) + (dy * dy)))
}
