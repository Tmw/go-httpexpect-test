package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestAPI(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(handler))
	e := httpexpect.Default(t, s.URL)

	body := e.GET("").
		Expect().
		Status(http.StatusOK).
		JSON()

	body.Path("$.items[*].id").Array().IsEqual([]int{1, 2})
	body.Path("$.items[*].color").Array().ConsistsOf("yellow", "pink")
	body.Path("$.items[*].weight").Array().ConsistsOf(12.12, 42.42)
}

func handler(w http.ResponseWriter,  r *http.Request) {
	w.Header().Add("content-type", "application/json")

	resp := map[string]any{
		"status": "ok",
		"items": []map[string]any {
			{
				"id": 1,
				"color": "yellow",
				"weight": 12.12,
			},
			{
				"id": 2,
				"color": "pink",
				"weight": 42.42,
			},
		},
	}

	json.NewEncoder(w).Encode(resp)
}