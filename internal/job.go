package internal

func NewCoord(x, y float64) *Coord {
	return &Coord{
		X: x,
		Y: y,
	}
}

type Job struct {
	ID             int
	Pickup         *Coord
	Dropoff        *Coord
	Cost           float64
	BaseTravelCost float64
}

func NewJob(id int, pickup, dropoff *Coord) *Job {
	return &Job{
		ID:             id,
		Pickup:         pickup,
		Dropoff:        dropoff,
		Cost:           CalculateCost(pickup, dropoff),
		BaseTravelCost: CalculateCost(NewCoord(0, 0), pickup),
	}
}

func (j *Job) CalculateTravelCost(start *Coord) float64 {
	return CalculateCost(start, j.Pickup)
}
