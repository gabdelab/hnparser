package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/1/count/", CountHandler)
    http.HandleFunc("/1/popular/", PopularHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
