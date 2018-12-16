package main

import (
    "fmt"
    "net/http"
    "strings"
)

func CountHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    date := strings.TrimPrefix(r.URL.Path, "/1/count/")

    if date == "" {
        // Empty date, returning a 400
        w.WriteHeader(http.StatusBadRequest)
        fmt.Println("no date given")
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Println("hello world")
}

func PopularHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    fmt.Println("not implemented")
}
