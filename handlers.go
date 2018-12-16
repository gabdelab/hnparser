package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Handler struct {
	logs logs
}

func GetHandler(logs logs) *Handler {
	return &Handler{logs: logs}
}

// Counter is the expected output for the count endpoint
type Counter struct {
	Count int `json:"count"`
}

// Query holds a query and a counter
type Query struct {
	Counter int    `json:"counter"`
	Query   string `json:"query"`
}

// Count implements the /count endpoint
func (h *Handler) Count(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	date := strings.TrimPrefix(r.URL.Path, "/1/count/")

	if date == "" {
		// Empty date, returning a 400
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("no date given")
		return
	}

	// Get the counter corresponding to the corresponding entry
	counter, err := getCounterFromEntry(h.logs, date)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(errors.Wrapf(err, "failed to get counter"))
		return
	}

	output := Counter{Count: counter}
	if err := json.NewEncoder(w).Encode(output); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(errors.Wrapf(err, "failed to properly encode json output"))
		return
	}
}

func (h *Handler) Popular(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	date := strings.TrimPrefix(r.URL.Path, "/1/popular/")
	if date == "" {
		// Empty date, returning a 400
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("no date given")
		return
	}

	// If no limit is set, use 0 as a limit, and return all results
	limitParam := "0"
	limits, ok := r.URL.Query()["limit"]
	if ok && len(limits[0]) == 1 {
		limitParam = limits[0]
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("invalid limit specified")
		return
	}

	// Get the popular queries corresponding to the corresponding entry
	popular, err := getTopQueriesFromEntry(h.logs, date, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(errors.Wrapf(err, "failed to get popular entries"))
		return
	}

	if err := json.NewEncoder(w).Encode(popular); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(errors.Wrapf(err, "failed to properly encode json output"))
		return
	}
}
