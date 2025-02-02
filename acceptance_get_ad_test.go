package main_test

import (
	"github.com/stretchr/testify/assert"
	"learning-go-challenges"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetAnAd(t *testing.T) {
	dateString := "2022-02-02 11:30:32"
	layout := "2006-01-02 15:04:05"
	currentTimestamp, _ := time.Parse(layout, dateString)

	dbConnection.Exec("INSERT INTO ads VALUES ('e85d27d4-3a6d-410f-a334-fdb52452fc17', 'A title', 'A description', 100, ?)", currentTimestamp)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ads/e85d27d4-3a6d-410f-a334-fdb52452fc17", nil)
	req.Header.Set("Content-Type", "application/json")

	main.HttpController.ServeHTTP(w, req)

	expectedResponse := `{
			"id": "e85d27d4-3a6d-410f-a334-fdb52452fc17" , 
			"title": "A title", 
			"description": "A description", 
			"price": 100, 
			"publishedAt": "2022-02-02 11:30:32 +0000 UTC"
		}`
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, expectedResponse, w.Body.String())
}

func TestGetNonExistingAd(t *testing.T) {
	dbConnection.Exec("TRUNCATE ads")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ads/e85d27d4-3a6d-410f-a334-fdb52452fc17", nil)
	req.Header.Set("Content-Type", "application/json")

	main.HttpController.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
