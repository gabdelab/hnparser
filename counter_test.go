package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addRoutesToList_works_on_empty_list(t *testing.T) {
	myMap := map[string]int{}
	routeLogs := map[string]int{"test": 2}
	addRoutesToList(myMap, routeLogs)
	assert.Equal(t, myMap, map[string]int{"test": 2})
}

func Test_addRoutesToList_can_add_two_elements(t *testing.T) {
	myMap := map[string]int{}
	routeLogs := map[string]int{"test": 2, "toast": 3}
	addRoutesToList(myMap, routeLogs)
	assert.Equal(t, len(myMap), 2)
}

func Test_addRoutesToList_doesnt_add_already_present_element(t *testing.T) {
	myMap := map[string]int{"route1": 2, "route2": 10}
	routeLogs := map[string]int{"test": 2, "route2": 3}
	addRoutesToList(myMap, routeLogs)
	assert.Equal(t, len(myMap), 3)
	assert.Equal(t, myMap["route2"], 13)
}

func Test_getCounterFromEntry_works_on_empty_logs(t *testing.T) {
	logs := logs{}
	total, err := getCounterFromEntry(logs, "2018-12")
	assert.Equal(t, 0, total)
	assert.Nil(t, err)
}

func Test_getCounterFromEntry_fails_on_invalid_date(t *testing.T) {
	logs := logs{}
	total, err := getCounterFromEntry(logs, "2018-12-")
	assert.Equal(t, 0, total)
	assert.NotNil(t, err)
}

var testLogs = logs{
	2018: {
		11: {
			12: {
				17: {
					52: {
						59: {
							"google.com": 3,
						},
					},
				},
				19: {
					11: {
						13: {
							"google.com": 7,
						},
						14: {
							"other": 1,
						},
					},
				},
			},
		},
	},
}

func Test_getCounterFromEntry_can_correctly_count_nominal_case(t *testing.T) {
	total, err := getCounterFromEntry(testLogs, "2018-11")
	assert.Equal(t, 2, total)
	assert.Nil(t, err)
	total, err = getCounterFromEntry(testLogs, "2018-11-12 19")
	assert.Equal(t, 2, total)
	assert.Nil(t, err)
	total, err = getCounterFromEntry(testLogs, "2018-11-12 19:11:13")
	assert.Equal(t, 1, total)
	assert.Nil(t, err)
}

func Test_getTopQueriesFromEntry_with_no_limit_returns_all_results(t *testing.T) {
	topQueries, err := getTopQueriesFromEntry(testLogs, "2018-11", 0)
	assert.Nil(t, err)
	assert.Equal(t, topQueries.Queries[0], Query{Counter: 10, Query: "google.com"})
	assert.Equal(t, topQueries.Queries[1], Query{Counter: 1, Query: "other"})
	assert.Equal(t, len(topQueries.Queries), 2)

}

func Test_getTopQueriesFromEntry_with_a_limit_returns_not_all_results(t *testing.T) {
	topQueries, err := getTopQueriesFromEntry(testLogs, "2018-11", 1)
	assert.Nil(t, err)
	assert.Equal(t, topQueries.Queries[0], Query{Counter: 10, Query: "google.com"})
	assert.Equal(t, len(topQueries.Queries), 1)
}

func Test_getTopQueriesFromEntry_with_higher_limit_returns_all_results(t *testing.T) {
	topQueries, err := getTopQueriesFromEntry(testLogs, "2018-11", 3)
	assert.Nil(t, err)
	assert.Equal(t, topQueries.Queries[0], Query{Counter: 10, Query: "google.com"})
	assert.Equal(t, topQueries.Queries[1], Query{Counter: 1, Query: "other"})
	assert.Equal(t, len(topQueries.Queries), 2)
}
