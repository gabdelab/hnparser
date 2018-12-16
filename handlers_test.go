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
	logs := logs{}
	handler := http.HandlerFunc(GetHandler(logs).Count)
	handler.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.JSONEq(t, `{"count":0}`, string(body))
}

func Test_CountHandler_returns_400_on_empty_date(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://localhost/1/count/", nil)
	logs := logs{}
	handler := http.HandlerFunc(GetHandler(logs).Count)
	handler.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, "", string(body))

}

func Test_PopularHandler_returns_200(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://localhost/1/popular/2018", nil)
	logs := logs{}
	handler := http.HandlerFunc(GetHandler(logs).Popular)
	handler.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "null\n", string(body))
}
