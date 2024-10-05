package internal

import (
	"fmt"
	"strings"
)

type Driver struct {
	Location      *Coord
	HomeLocation  *Coord
	WorkCompleted float64
	CompletedJobs []*Job

	maxAllowableWork float64
}

func NewDriver(startingLocation *Coord, maxWorkdAllowable float64) *Driver {
	return &Driver{
		Location:         startingLocation,
		HomeLocation:     startingLocation,
		maxAllowableWork: maxWorkdAllowable,
	}
}

func (d *Driver) AvailableForJob(job *Job) (bool, float64) {
	// check for job feasiblity
	if d.WorkCompleted+job.Cost > d.maxAllowableWork {
		return false, -1
	}

	distanceToJob := CalculateCost(d.Location, job.Pickup)
	distanceHome := CalculateCost(job.Dropoff, d.HomeLocation)

	remainingWork := d.maxAllowableWork - (d.WorkCompleted + distanceToJob + job.Cost + distanceHome)

	return remainingWork >= 0, distanceToJob
}

func (d *Driver) AssignJob(job *Job) {
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
