package main

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
)

func Test_CountHandler_returns_200(t *testing.T) {
    w := httptest.NewRecorder()
    r := httptest.NewRequest("GET", "http://localhost/1/count/2018", nil)

    handler := http.HandlerFunc(CountHandler)
    handler.ServeHTTP(w, r)

    resp := w.Result()
    body, _ := ioutil.ReadAll(resp.Body)
    assert.Equal(t, http.StatusOK, resp.StatusCode)
    assert.Equal(t, "", string(body))
}

func Test_CountHandler_returns_400_on_empty_date(t *testing.T) {
    w := httptest.NewRecorder()
    r := httptest.NewRequest("GET", "http://localhost/1/count/", nil)
    handler := http.HandlerFunc(CountHandler)
    handler.ServeHTTP(w, r)

    resp := w.Result()
    body, _ := ioutil.ReadAll(resp.Body)
    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
    assert.Equal(t, "", string(body))

}

func Test_PopularHandler_returns_200(t *testing.T) {
    w := httptest.NewRecorder()
    r := httptest.NewRequest("GET", "http://localhost/1/popular/2018", nil)
    handler := http.HandlerFunc(PopularHandler)
    handler.ServeHTTP(w, r)

    resp := w.Result()
    body, _ := ioutil.ReadAll(resp.Body)
    assert.Equal(t, http.StatusOK, resp.StatusCode)
    assert.Equal(t, "", string(body))
}
