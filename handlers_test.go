package main

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
)

func Test_CountHandler_returns_200(t *testing.T) {
    w := httptest.NewRecorder()
    r := httptest.NewRequest("GET", "http://localhost/1/count/2018", nil)

    handler := http.HandlerFunc(CountHandler)
    handler.ServeHTTP(w, r)

    resp := w.Result()
    body, _ := ioutil.ReadAll(resp.Body)
    if resp.StatusCode != 200 {
        t.Error("expected a 200 status code")
    }
    if string(body) != "" {
        t.Error("expected an empty body")
    }
}

func Test_CountHandler_returns_400_on_empty_date(t *testing.T) {
    w := httptest.NewRecorder()
    r := httptest.NewRequest("GET", "http://localhost/1/count/", nil)
    handler := http.HandlerFunc(CountHandler)
    handler.ServeHTTP(w, r)

    resp := w.Result()
    body, _ := ioutil.ReadAll(resp.Body)
    if resp.StatusCode != 400 {
        t.Errorf("expected a 400 status code, got %v", resp.StatusCode)
    }
    if string(body) != "" {
        t.Error("expected an empty body")
    }
}

func Test_PopularHandler_returns_200(t *testing.T) {
    w := httptest.NewRecorder()
    r := httptest.NewRequest("GET", "http://localhost/1/popular/2018", nil)
    handler := http.HandlerFunc(PopularHandler)
    handler.ServeHTTP(w, r)

    resp := w.Result()
    body, _ := ioutil.ReadAll(resp.Body)
    if resp.StatusCode != 200 {
        t.Error("expected a 200 status code")
    }
    if string(body) != "" {
        t.Error("expected an empty body")
    }
}
