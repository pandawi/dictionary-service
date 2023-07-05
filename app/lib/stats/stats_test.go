package stats_test

import (
	"dictionary-service/app/lib/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatsImpl_CountWords(t *testing.T) {
	s := stats.NewStats()
	length := s.CountWords([]string{"hi", "panda"})
	assert.Equal(t, 2, length)
}

func TestStatsImpl_CalculateAvgReqProcessingTime(t *testing.T) {
	s := stats.NewStats()
	s.AddRequest(1)
	s.AddRequest(2)
	s.AddRequest(3)
	totalTime := s.CalculateAvgReqProcessingTime()
	assert.Equal(t, 2, totalTime)
}
