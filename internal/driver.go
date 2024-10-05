package internal

import (
	"fmt"
	"strings"
)

type Driver struct {
	ID            int
	Location      *Coord
	HomeLocation  *Coord
	WorkCompleted float64
	CompletedJobs []*Job

	maxAllowableWork float64
	AtHome           bool
}

func NewDriver(id int, startingLocation *Coord, maxWorkdAllowable float64) *Driver {
	return &Driver{
		ID:               id,
		Location:         startingLocation,
		HomeLocation:     NewCoord(0, 0),
		maxAllowableWork: maxWorkdAllowable,
	}
}

func (d *Driver) AvailableForJob(job *Job) (bool, float64) {
	if d.WorkCompleted+job.Cost > d.maxAllowableWork {
		return false, -1
	}

	distanceToJob := CalculateCost(d.Location, job.Pickup)
	distanceForJob := job.Cost
	distanceHome := CalculateCost(job.Dropoff, d.HomeLocation)

	remainingWork := d.maxAllowableWork - (d.WorkCompleted + distanceToJob + distanceForJob + distanceHome)

	return remainingWork >= 0, distanceToJob
}

func (d *Driver) AssignJob(job *Job) {
	if d.AtHome {
		d.AtHome = false
	}

	distanceToJob := CalculateCost(d.Location, job.Pickup)
	d.WorkCompleted = d.WorkCompleted + distanceToJob + job.Cost
	d.Location = job.Dropoff
	d.CompletedJobs = append(d.CompletedJobs, job)
}

func (d *Driver) simpleString() string {
	jobIDs := []string{}
	for _, job := range d.CompletedJobs {
		jobIDs = append(jobIDs, fmt.Sprintf("%d", job.ID))
	}

	return fmt.Sprintf("[%s]", strings.Join(jobIDs, ", "))
}

func (d *Driver) String() string {
	return d.simpleString()
}
