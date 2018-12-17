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
	assert.JSONEq(t, `{"queries":null}`, string(body))
}

func Test_PopularHandler_returns_400_on_empty_date(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://localhost/1/popular/", nil)
	logs := logs{}
	handler := http.HandlerFunc(GetHandler(logs).Popular)
	handler.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, "", string(body))
}

var exampleLogs = logs{
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

func Test_CountHandler_returns_200_with_counter(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://localhost/1/count/2018", nil)
	handler := http.HandlerFunc(GetHandler(exampleLogs).Count)
	handler.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.JSONEq(t, `{"count":2}`, string(body))
}

func Test_PopularHandler_returns_200_with_results_limited(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://localhost/1/popular/2018?limit=1", nil)
	handler := http.HandlerFunc(GetHandler(exampleLogs).Popular)
	handler.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.JSONEq(t, `{"queries": [{"counter":10, "query": "google.com"}]}`, string(body))
}
