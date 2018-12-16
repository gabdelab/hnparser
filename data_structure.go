package main

// We choose a nested data structure, by date
//
// This is because we'll always have to filter by date.
// So, searching in the whole struct will be quite easy
// and we'll quickly reach small amounts of data that
// are easy to parse.
//
// The different level of keys are:
// year - month - day - hour - minute - second - url
// And then there's only the counter
//
// If we had to implement another kind of filter, like
// for instance based on the route, this wouldn't be efficient
// anymore, so we would have to use another indexing

type routelogs map[string]int
type secondlogs map[int]routelogs
type minutelogs map[int]secondlogs
type hourlogs map[int]minutelogs
type daylogs map[int]hourlogs
type monthlogs map[int]daylogs
type logs map[int]monthlogs
