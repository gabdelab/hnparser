package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/pkg/errors"
)

// addRouteToList adds all urls in routeLogs to myMap if not present
//
// to achieve this fastly, we use maps with O(1) access
func addRoutesToList(myMap map[string]int, routeLogs routelogs) {
	for route, counter := range routeLogs {
		_, ok := myMap[route]
		if !ok {
			myMap[route] = counter
		} else {
			myMap[route] += counter
		}
	}
}

func addRoutesFromSecondsToList(myMap map[string]int, secondLogs secondlogs) {
	for _, routes := range secondLogs {
		addRoutesToList(myMap, routes)
	}
	return
}

func addRoutesFromMinutesToList(myMap map[string]int, minuteLogs minutelogs) {
	for _, seconds := range minuteLogs {
		addRoutesFromSecondsToList(myMap, seconds)
	}
	return
}

func addRoutesFromHoursToList(myMap map[string]int, hourLogs hourlogs) {
	for _, minutes := range hourLogs {
		addRoutesFromMinutesToList(myMap, minutes)
	}
	return
}

func addRoutesFromDaysToList(myMap map[string]int, dayLogs daylogs) {
	for _, hours := range dayLogs {
		addRoutesFromHoursToList(myMap, hours)
	}
	return
}

func addRoutesFromMonthsToList(myMap map[string]int, monthLogs monthlogs) {
	for _, days := range monthLogs {
		addRoutesFromDaysToList(myMap, days)
	}
	return
}

var (
	withSeconds = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`)
	withMinutes = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}`)
	withHours   = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}`)
	withDays    = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}`)
	withMonths  = regexp.MustCompile(`^\d{4}-\d{2}`)
	withYears   = regexp.MustCompile(`^\d{4}`)
)

// getMatchingRoutes returns the number of different entries
// belonging to a given date
//
// It has two main purposes:
// - find out what level of struct we have to inspect, by parsing the parameter
// - create the subset of entries that will be counted
func getMatchingRoutes(logs logs, date string) (map[string]int, error) {
	urls := map[string]int{}

	switch {
	case withSeconds.FindString(date) == date:
		// If we match the whole date, this is the simplest case;
		// We just have to count the entries
		year, _ := strconv.Atoi(date[0:4])
		month, _ := strconv.Atoi(date[5:7])
		day, _ := strconv.Atoi(date[8:10])
		hour, _ := strconv.Atoi(date[11:13])
		minute, _ := strconv.Atoi(date[14:16])
		second, _ := strconv.Atoi(date[17:19])
		val, ok := logs[year][month][day][hour][minute][second]
		if !ok {
			fmt.Println("not found, count is 0")
			return urls, nil
		}
		return val, nil
	case withMinutes.FindString(date) == date:
		year, _ := strconv.Atoi(date[0:4])
		month, _ := strconv.Atoi(date[5:7])
		day, _ := strconv.Atoi(date[8:10])
		hour, _ := strconv.Atoi(date[11:13])
		minute, _ := strconv.Atoi(date[14:16])
		val, ok := logs[year][month][day][hour][minute]
		if !ok {
			fmt.Println("not found, count is 0")
			return urls, nil
		}
		addRoutesFromSecondsToList(urls, val)
		return urls, nil

	case withHours.FindString(date) == date:
		year, _ := strconv.Atoi(date[0:4])
		month, _ := strconv.Atoi(date[5:7])
		day, _ := strconv.Atoi(date[8:10])
		hour, _ := strconv.Atoi(date[11:13])
		val, ok := logs[year][month][day][hour]
		if !ok {
			fmt.Println("not found, count is 0")
			return urls, nil
		}
		addRoutesFromMinutesToList(urls, val)
		return urls, nil

	case withDays.FindString(date) == date:
		year, _ := strconv.Atoi(date[0:4])
		month, _ := strconv.Atoi(date[5:7])
		day, _ := strconv.Atoi(date[8:10])
		val, ok := logs[year][month][day]
		if !ok {
			fmt.Println("not found, count is 0")
			return urls, nil
		}
		addRoutesFromHoursToList(urls, val)
		return urls, nil

	case withMonths.FindString(date) == date:
		year, _ := strconv.Atoi(date[0:4])
		month, _ := strconv.Atoi(date[5:7])
		val, ok := logs[year][month]
		if !ok {
			fmt.Println("not found, count is 0")
			return urls, nil
		}
		addRoutesFromDaysToList(urls, val)
		return urls, nil

	case withYears.FindString(date) == date:
		year, _ := strconv.Atoi(date[0:4])
		val, ok := logs[year]
		if !ok {
			fmt.Println("not found, count is 0")
			return urls, nil
		}
		addRoutesFromMonthsToList(urls, val)
		return urls, nil

	default:
		return urls, errors.New("invalid date format")
	}
}

func getCounterFromEntry(logs logs, date string) (int, error) {
	matchingRoutes, err := getMatchingRoutes(logs, date)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get counter from entry")
	}
	return len(matchingRoutes), nil
}

func getTopQueriesFromEntry(logs logs, date string, limit int) ([]Query, error) {
	matchingRoutes, err := getMatchingRoutes(logs, date)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get counter from entry")
	}

	// The results need to be sorted by counter, so that the "limit" parameter filters out
	// correctly the least frequent requests.
	var ss []Query
	for key, value := range matchingRoutes {
		ss = append(ss, Query{Query: key, Counter: value})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Counter > ss[j].Counter
	})

	if limit == 0 {
		return ss, nil
	}

	if limit > len(ss) {
		return ss, nil
	}

	return ss[0:limit], nil
}
