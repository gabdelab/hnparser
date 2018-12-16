package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

// GetCounterFromEntry returns the number of different entries
// belonging to a given date
func (h *Handler) GetCounterFromEntry(date string) (int, error) {
	// Parse date : how many parameters do we have ?

	urls := []string{}
	// Find the subtree of logs

	// Create a reverse list with all entries, no matter their counter

	// Return the number
	return len(urls), nil
}

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
	counter, err := h.GetCounterFromEntry(date)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(errors.Wrapf(err, "failed to get counter"))
		return
	}

	output := Counter{Count: counter}
	if err := json.NewEncoder(w).Encode(output); err != nil {
		fmt.Println("here")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(errors.Wrapf(err, "failed to properly encode json output"))
		return
	}
}

func (h *Handler) Popular(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Println("not implemented")
}
