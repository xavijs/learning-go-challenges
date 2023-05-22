package main_test

import (
	"bytes"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"learning-go-challenges"
	"learning-go-challenges/domain/ad"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestPostAnAd(t *testing.T) {
	router := main.SetupHttpRouter()

	w := httptest.NewRecorder()

	requestBody := `{"title": "A Title", "description": "An ad description", "price": 100}`
	req, _ := http.NewRequest("POST", "/ads", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	postedAdId := extractIdFromJsonResponse(w.Body.String())

	assert.Equal(t, http.StatusCreated, w.Code)
	verifyAdExistsInDb(postedAdId, t)
}

func TestFailWhenMissingDataInRequest(t *testing.T) {
	*main.RepositoryMemory = []ad.Ad{}
	router := main.SetupHttpRouter()

	w := httptest.NewRecorder()

	requestBody := `{"title": "A Title"}`
	req, _ := http.NewRequest("POST", "/ads", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Empty(t, main.RepositoryMemory)
}

func verifyAdExistsInDb(id string, t *testing.T) {
	var foundAd ad.Ad
	for _, ad := range *main.RepositoryMemory {
		if ad.Id.Value == id {
			foundAd = ad
		}
	}

	expectedAd := ad.Ad{
		Id:          ad.Id{Value: id},
		Title:       "A Title",
		Description: "An ad description",
		Price:       100,
		PublishedAt: time.Now().UTC(),
	}

	assert.Equal(t, expectedAd.Price, foundAd.Price)
	assert.Equal(t, expectedAd.Title, foundAd.Title)
	assert.Equal(t, expectedAd.Description, foundAd.Description)
	assertTimeCloseToNowUTC(foundAd.PublishedAt, t)
}

func extractIdFromJsonResponse(jsonResponse string) string {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonResponse), &data)
	if err != nil {
		return "Cannot extract id from response: " + jsonResponse
	}
	return data["id"].(string)
}

func assertTimeCloseToNowUTC(timeToVerify time.Time, t *testing.T) {
	currentTime := time.Now().UTC()
	threshold := 10 * time.Second
	diff := currentTime.Sub(timeToVerify)
	isNear := diff < threshold && diff > -threshold

	assert.True(t, isNear, "Time "+timeToVerify.String()+"is too different than now.")
}
