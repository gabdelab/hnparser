package main

import (
    "fmt"
    "regexp"
    "strconv"

    "github.com/pkg/errors"
)

// addRouteToList adds all urls in routeLogs to myList if not present
func addRoutesToList(myList []string, routeLogs routelogs) []string {
    for route := range routeLogs {
        isPresent := false
        for _, registeredRoute := range myList {
            if route == registeredRoute {
                isPresent = true
                break
            }
        }
        if !isPresent {
            myList = append(myList, route)
        }
    }
    return myList
}

func addRoutesFromSecondsToList(myList []string, secondLogs secondlogs) []string {
    for _, routes := range secondLogs {
        myList = addRoutesToList(myList, routes)
    }
    return myList
}

func addRoutesFromMinutesToList(myList []string, minuteLogs minutelogs) []string {
    for _, seconds := range minuteLogs {
        myList = addRoutesFromSecondsToList(myList, seconds)
    }
    return myList
}

func addRoutesFromHoursToList(myList []string, hourLogs hourlogs) []string {
    for _, minutes := range hourLogs {
        myList = addRoutesFromMinutesToList(myList, minutes)
    }
    return myList
}

func addRoutesFromDaysToList(myList []string, dayLogs daylogs) []string {
    for _, hours := range dayLogs {
        myList = addRoutesFromHoursToList(myList, hours)
    }
    return myList
}

func addRoutesFromMonthsToList(myList []string, monthLogs monthlogs) []string {
    for _, days := range monthLogs {
        myList = addRoutesFromDaysToList(myList, days)
    }
    return myList
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
    urls := []string{}

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
        urls = addRoutesFromSecondsToList(urls, val)
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
        urls = addRoutesFromMinutesToList(urls, val)
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
        urls = addRoutesFromHoursToList(urls, val)
        return len(urls), nil

    case withMonths.FindString(date) == date:
        year, _ := strconv.Atoi(date[0:4])
        month, _ := strconv.Atoi(date[5:7])
        val, ok := logs[year][month]
        if !ok {
            fmt.Println("not found, count is 0")
            return 0, nil
        }
        urls = addRoutesFromDaysToList(urls, val)
        return len(urls), nil

    case withYears.FindString(date) == date:
        year, _ := strconv.Atoi(date[0:4])
        val, ok := logs[year]
        if !ok {
            fmt.Println("not found, count is 0")
            return 0, nil
        }
        urls = addRoutesFromMonthsToList(urls, val)
        return len(urls), nil

    default:
        return 0, errors.New("invalid date format")
    }
}
