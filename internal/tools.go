package internal

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

func LoadProblem(path string) ([]*Job, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	jobs := []*Job{}
	for i, row := range rows {
		if i == 0 {
			continue
		}

		// ID (0,0) (0,0)
		id, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, err
		}

		pickupX, pickupY, err := parseStringCoord(row[1])
		if err != nil {
			return nil, err
		}

		dropOffX, dropOffY, err := parseStringCoord(row[2])
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, NewJob(
			id,
			NewCoord(pickupX, pickupY),
			NewCoord(dropOffX, dropOffY),
		))
	}

	return jobs, nil
}

// parses a string (0,0) to a pair of floats
func parseStringCoord(input string) (x, y float64, err error) {
	input = strings.TrimPrefix(input, "(")
	input = strings.TrimSuffix(input, ")")
	splitValues := strings.Split(input, ",")

	x, err = strconv.ParseFloat(splitValues[0], 64)
	if err != nil {
		return
	}

	y, err = strconv.ParseFloat(splitValues[1], 64)
	if err != nil {
		return
	}

	return
}
