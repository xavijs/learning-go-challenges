package main_test

import (
	"bytes"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"learning-go-challenges"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestPostAnAd(t *testing.T) {
	w := httptest.NewRecorder()

	requestBody := `{"title": "A Title", "description": "An ad description", "price": 100}`
	req, _ := http.NewRequest("POST", "/ads", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")

	main.HttpController.ServeHTTP(w, req)
	postedAdId := extractIdFromJsonResponse(w.Body.String())

	assert.Equal(t, http.StatusCreated, w.Code)
	verifyAdExistsInDb(postedAdId, t)
}

func TestFailWhenMissingDataInRequest(t *testing.T) {
	w := httptest.NewRecorder()

	requestBody := `{"title": "A Title"}`
	req, _ := http.NewRequest("POST", "/ads", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")

	main.HttpController.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func verifyAdExistsInDb(id string, t *testing.T) {
	var rows []map[string]interface{}
	sql := fmt.Sprintf("SELECT id, title, description, price, published_at FROM ads WHERE id = '%v';", id)
	dbConnection.Raw(sql).Scan(&rows)
	dbAd := rows[0]
	assert.EqualValues(t, 100, dbAd["price"])
	assert.EqualValues(t, "A Title", dbAd["title"])
	assert.EqualValues(t, "An ad description", dbAd["description"])
	assertTimeCloseToNowUTC(t, dbAd["published_at"].(time.Time))
}

func extractIdFromJsonResponse(jsonResponse string) string {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonResponse), &data)
	if err != nil {
		return "Cannot extract id from response: " + jsonResponse
	}
	return data["id"].(string)
}

func assertTimeCloseToNowUTC(t *testing.T, timeToVerify time.Time) {
	currentTime := time.Now().UTC()
	threshold := 10 * time.Second
	diff := currentTime.Sub(timeToVerify)
	isNear := diff < threshold && diff > -threshold

	assert.True(t, isNear, "Time "+timeToVerify.String()+"is too different than now.")
}
