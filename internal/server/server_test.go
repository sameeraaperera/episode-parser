package server_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sap4001/episode-parser/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestDefaultRoute(t *testing.T) {
	assert.True(t, true)

	router := server.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Welcome to episode-parser")
}

func TestParseRouteValidJSON(t *testing.T) {
	var inputJSON = []byte(`{
		"payload": [
			{
				"country": "UK",
				"description": "What's life like when you have enough children to field your own football team?",
				"drm": true,
				"episodeCount": 3,
				"genre": "Reality",
				"image": {
					"showImage": "http://catchup.ninemsn.com.au/img/jump-in/shows/16KidsandCounting1280.jpg"
				},
				"language": "English",
				"nextEpisode": null,
				"primaryColour": "#ff7800",
				"seasons": [
					{
						"slug": "show/16kidsandcounting/season/1"
					}
				],
				"slug": "show/16kidsandcounting",
				"title": "16 Kids and Counting",
				"tvChannel": "GEM"
			}
		],
		"skip": 0,
		"take": 10,
		"totalRecords": 75
	}`)
	wantStr := `{"Response":[{"Image":"http://catchup.ninemsn.com.au/img/jump-in/shows/16KidsandCounting1280.jpg","Slug":"show/16kidsandcounting","Title":"16 Kids and Counting"}]}`

	router := server.SetupRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/parse", bytes.NewBuffer(inputJSON))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), wantStr)
}

func TestParseRouteInValidJSON(t *testing.T) {
	var inputJSON = []byte(`{
		"invalid":true
	}`)
	wantStr := `{"error":"Could not decode request: JSON parsing failed"}`

	router := server.SetupRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/parse", bytes.NewBuffer(inputJSON))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, w.Body.String(), wantStr)
}
