package main

import (
	"dictionary-service/app"
	"dictionary-service/app/data/cache"
	"dictionary-service/app/data/db"
	"dictionary-service/app/lib/stats"
	"dictionary-service/app/lib/trie"
	"dictionary-service/app/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ServiceSearch(t *testing.T) {
	mockTrie := trie.NewTrie()
	mockData := []string{"adnap", "aadnp"}
	mockTrie.FromSource(mockData)
	mockCache := cache.NewCache(1)
	s := &stats.StatsImpl{}
	mockApp := &app.App{Trie: &mockTrie, Stats: s, Cache: mockCache}
	req, _ := http.NewRequest("GET", "/api/v1/similar?word=panda", nil)
	w := httptest.NewRecorder()
	r := setupRouter(mockApp)
	r.ServeHTTP(w, req)
	body := w.Body.String()
	var response models.SearchResponse
	jsonErr := json.Unmarshal([]byte(body), &response)
	assert.Equal(t, nil, jsonErr)
	assert.Equal(t, 200, w.Code)
	assert.ElementsMatch(t, response.Similar, mockData)
}

func Test_ServiceStats(t *testing.T) {
	s := stats.NewStats()
	mockTrie := trie.NewTrie()
	mockDb := db.NewDb()
	mockCache := cache.NewCache(1)
	mockApp := &app.App{Stats: s, Trie: &mockTrie, DB: mockDb, Cache: mockCache}
	searchReq, _ := http.NewRequest("GET", "/api/v1/similar?word=panda", nil)
	statsReq, _ := http.NewRequest("GET", "/api/v1/stats", nil)
	responses := setup(mockApp, []*http.Request{searchReq, statsReq})
	sr := responses[1]
	body := sr.Body.String()
	var response models.StatsResponse
	jsonErr := json.Unmarshal([]byte(body), &response)
	assert.Equal(t, nil, jsonErr)
	assert.Equal(t, 200, sr.Code)
	assert.Equal(t, response.TotalRequests, 1)
	assert.NotZero(t, response.AvgProcessingTimeNs)
	assert.Equal(t, response.TotalWords, 0)
}

func setup(mockApp *app.App, reqs []*http.Request) []*httptest.ResponseRecorder {
	r := setupRouter(mockApp)
	recorders := make([]*httptest.ResponseRecorder, len(reqs))
	for i, req := range reqs {
		w := httptest.NewRecorder()
		recorders[i] = w
		r.ServeHTTP(w, req)
	}
	return recorders
}
