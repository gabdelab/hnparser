package main

import (
    "fmt"
    "regexp"
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

// getCounterFromEntry returns the number of different entries
// belonging to a given date
//
// It has two main purposes:
// - find out what level of struct we have to inspect, by parsing the parameter
// - count the number of entries
func getCounterFromEntry(logs logs, date string) (int, error) {
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
            return 0, nil
        }
        return len(val), nil
    case withMinutes.FindString(date) == date:
        year, _ := strconv.Atoi(date[0:4])
        month, _ := strconv.Atoi(date[5:7])
        day, _ := strconv.Atoi(date[8:10])
        hour, _ := strconv.Atoi(date[11:13])
        minute, _ := strconv.Atoi(date[14:16])
        val, ok := logs[year][month][day][hour][minute]
        if !ok {
            fmt.Println("not found, count is 0")
            return 0, nil
        }
        addRoutesFromSecondsToList(urls, val)
        return len(urls), nil

    case withHours.FindString(date) == date:
        year, _ := strconv.Atoi(date[0:4])
        month, _ := strconv.Atoi(date[5:7])
        day, _ := strconv.Atoi(date[8:10])
        hour, _ := strconv.Atoi(date[11:13])
        val, ok := logs[year][month][day][hour]
        if !ok {
            fmt.Println("not found, count is 0")
            return 0, nil
        }
        addRoutesFromMinutesToList(urls, val)
        return len(urls), nil

    case withDays.FindString(date) == date:
        year, _ := strconv.Atoi(date[0:4])
        month, _ := strconv.Atoi(date[5:7])
        day, _ := strconv.Atoi(date[8:10])
        val, ok := logs[year][month][day]
        if !ok {
            fmt.Println("not found, count is 0")
            return 0, nil
        }
        addRoutesFromHoursToList(urls, val)
        return len(urls), nil

    case withMonths.FindString(date) == date:
        year, _ := strconv.Atoi(date[0:4])
        month, _ := strconv.Atoi(date[5:7])
        val, ok := logs[year][month]
        if !ok {
            fmt.Println("not found, count is 0")
            return 0, nil
        }
        addRoutesFromDaysToList(urls, val)
        return len(urls), nil

    case withYears.FindString(date) == date:
        year, _ := strconv.Atoi(date[0:4])
        val, ok := logs[year]
        if !ok {
            fmt.Println("not found, count is 0")
            return 0, nil
        }
        addRoutesFromMonthsToList(urls, val)
        return len(urls), nil

    default:
        return 0, errors.New("invalid date format")
    }
}
