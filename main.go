package main

import (
	"deliver-driver-problem/internal"
	"fmt"
	"os"
)

func main() {
	path := os.Args[1]

	jobs, err := internal.LoadProblem(path)
	if err != nil {
		panic(fmt.Sprintf("failed to load job list: %s", err))
	}

	dispatcher := internal.NewDispatcher(internal.MAX_TRAVEL_DISTANCE, internal.NEW_DRIVER_COST)
	dispatcher.Dispatch(jobs)
}
