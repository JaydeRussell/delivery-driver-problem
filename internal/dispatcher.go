package internal

import (
	"fmt"
	"math"
)

type Dispatcher struct {
	drivers           []*Driver
	maxWorkdAllowable float64
}

func NewDispatcher(maxAllowableWork float64) *Dispatcher {
	return &Dispatcher{
		maxWorkdAllowable: maxAllowableWork,
		drivers:           []*Driver{},
	}
}

func (d *Dispatcher) Dispatch(jobs []*Job) {
	//sortedJobs := jobs //d.groupJobs(jobs)
	setSize := len(jobs)
	currentDriver := d.createNewDriver()

	for i := 0; i < setSize; i++ {
		var closestJob *Job
		lowestCost := math.MaxFloat64
		index := 0
		for j, job := range jobs {
			if ok, driverTravelCost := currentDriver.AvailableForJob(job); ok &&
				(driverTravelCost < lowestCost) {
				closestJob = job
				lowestCost = driverTravelCost
				index = j
			}
		}
		if closestJob != nil {
			currentDriver.AssignJob(closestJob)
			jobs = append(jobs[:index], jobs[index+1:]...)
		} else {
			currentDriver = d.createNewDriver()
			i--
		}
	}

	d.printResults()
}

func (d *Dispatcher) createNewDriver() *Driver {
	newDriver := NewDriver(0, NewCoord(0, 0), d.maxWorkdAllowable)
	d.drivers = append(d.drivers, newDriver)
	return newDriver
}

func (d *Dispatcher) printResults() {
	for _, driver := range d.drivers {
		if len(driver.CompletedJobs) > 0 {
			fmt.Println(driver)
		}
	}
}
