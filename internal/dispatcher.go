package internal

import (
	"fmt"
	"math"
)

type Dispatcher struct {
	drivers           []*Driver
	maxWorkdAllowable float64
	newDriverCost     float64
}

func NewDispatcher(maxAllowableWork float64, newDriverCost float64) *Dispatcher {
	return &Dispatcher{
		maxWorkdAllowable: maxAllowableWork,
		newDriverCost:     newDriverCost,
		drivers:           []*Driver{},
	}
}

func (d *Dispatcher) Dispatch(jobs []*Job) {
	currentDriver := d.createNewDriver()

	setSize := len(jobs)
	for i := 0; i < setSize; i++ {
		var closestJob *Job
		lowestCost := math.MaxFloat64
		index := 0
		// find the nearest job to our driver
		for j, job := range jobs {
			if ok, driverTravelCost := currentDriver.AvailableForJob(job); ok &&
				driverTravelCost < job.BaseTravelCost+d.newDriverCost &&
				driverTravelCost < lowestCost {
				closestJob = job
				lowestCost = driverTravelCost
				index = j
			}
		}
		// if we found a job, assign it
		if closestJob != nil {
			currentDriver.AssignJob(closestJob)
			jobs = append(jobs[:index], jobs[index+1:]...)
		} else {
			// no jobs were found, so we create a new driver and start over
			currentDriver = d.createNewDriver()
			i--
		}
	}

	d.printResults()
}

func (d *Dispatcher) createNewDriver() *Driver {
	newDriver := NewDriver(NewCoord(0, 0), d.maxWorkdAllowable)
	d.drivers = append(d.drivers, newDriver)
	return newDriver
}

func (d *Dispatcher) printResults() {
	for _, driver := range d.drivers {
		fmt.Println(driver)
	}
}
