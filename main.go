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

	dispatcher := internal.NewDispatcher(60 * 12)
	dispatcher.Dispatch(jobs)
}
