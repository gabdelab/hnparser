package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// parseLogs parses a file given by its path and fills the data in-memory
func parseLogs(filePath string, output logs) error {
	if filePath == "" {
		return errors.New("missing log file")
	}

	// Open file and read it
	file, err := os.Open(filePath)
	if err != nil {
		return errors.Wrap(err, "could not open logs file")
	}
	defer file.Close()

	// Use a csv reader with space as delimitor
	reader := csv.NewReader(file)
	reader.Comma = '\t'

	data, err := reader.ReadAll()
	if err != nil {
		return errors.Wrap(err, "failed to read tsv data")
	}

	for _, line := range data {
		err := processLine(output, line)
		if err != nil {
			return errors.Wrap(err, "at least one line is wrong, the file looks invalid")
		}
	}
	return nil
}

// processLine processes a line from the tsv
//
// it is expected to look as is:
// ["yyyy-mm-dd HH:mm:ss", "url"]
// If the line is misformatted, this raises an error
func processLine(logs logs, line []string) error {
	if len(line) != 2 {
		return errors.Errorf("invalid line %s %d", line, len(line))
	}
	// Parse date and time at once
	// Expected dates are ISO8601, just convert to RFC 3339 for easier parsing
	splittedTime := strings.Split(line[0], " ")
	datetime := fmt.Sprintf("%sT%sZ", splittedTime[0], splittedTime[1])
	t, err := time.Parse(time.RFC3339, datetime)
	if err != nil {
		return errors.Errorf("invalid time %s", datetime)
	}

	// Add the corresponding entry in the map if it doesn't exist yeat
	if _, ok := logs[t.Year()]; !ok {
		logs[t.Year()] = monthlogs{}
	}
	if _, ok := logs[t.Year()][int(t.Month())]; !ok {
		logs[t.Year()][int(t.Month())] = daylogs{}
	}
	if _, ok := logs[t.Year()][int(t.Month())][t.Day()]; !ok {
		logs[t.Year()][int(t.Month())][t.Day()] = hourlogs{}
	}
	if _, ok := logs[t.Year()][int(t.Month())][t.Day()][t.Hour()]; !ok {
		logs[t.Year()][int(t.Month())][t.Day()][t.Hour()] = minutelogs{}
	}
	if _, ok := logs[t.Year()][int(t.Month())][t.Day()][t.Hour()][t.Minute()]; !ok {
		logs[t.Year()][int(t.Month())][t.Day()][t.Hour()][t.Minute()] = secondlogs{}
	}
	if _, ok := logs[t.Year()][int(t.Month())][t.Day()][t.Hour()][t.Minute()][t.Second()]; !ok {
		logs[t.Year()][int(t.Month())][t.Day()][t.Hour()][t.Minute()][t.Second()] = routelogs{}
	}

	logs[t.Year()][int(t.Month())][t.Day()][t.Hour()][t.Minute()][t.Second()][line[1]] += 1
	return nil
}
