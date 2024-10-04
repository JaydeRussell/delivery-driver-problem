package internal

import (
	"fmt"
	"math"
	"slices"
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
	sortedJobs := d.groupJobs(jobs)

	for _, job := range sortedJobs {
		var idealDriver *Driver
		idealDriverCost := CalculateCost(NewCoord(0, 0), job.Pickup)

		for _, driver := range d.drivers {
			if available, cost := driver.AvailableForJob(job); available {
				if cost < idealDriverCost {
					idealDriver = driver
					idealDriverCost = cost
				}
			}
		}

		if idealDriver == nil {
			d.assignNewDriver(job)
		} else {
			idealDriver.AssignJob(job)
		}
	}

	for _, driver := range d.drivers {
		fmt.Println(driver)
	}
}

func (d *Dispatcher) groupJobs(jobs []*Job) []*Job {
	newJobList := []*Job{}

	currentNode := NewCoord(0, 0)
	setSize := len(jobs)
	for i := 0; i < setSize; i++ {
		var cheapestJob *Job
		cheapestCost := math.MaxFloat64
		index := 0
		for i, job := range jobs {
			if cost := CalculateCost(currentNode, job.Pickup); cost < cheapestCost {
				cheapestJob = job
				cheapestCost = cost
				index = i
			}
		}

		newJobList = append(newJobList, cheapestJob)
		currentNode = cheapestJob.Dropoff
		jobs = append(jobs[:index], jobs[index+1:]...)
	}
	slices.Reverse[[]*Job](newJobList)

	return newJobList
}

func (d *Dispatcher) assignNewDriver(job *Job) {
	newDriver := NewDriver(len(d.drivers), NewCoord(0, 0), d.maxWorkdAllowable)
	if available, _ := newDriver.AvailableForJob(job); !available {
		panic(fmt.Sprintf("job: %+v could not be assigned", *job))
	}

	newDriver.AssignJob(job)

	d.drivers = append(d.drivers, newDriver)
}
