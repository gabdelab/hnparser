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
	w.WriteHeader(http.StatusOK)
	fmt.Println("not implemented")
}
