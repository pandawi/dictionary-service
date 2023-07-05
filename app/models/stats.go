package models

type StatsResponse struct {
	TotalWords          int `json:"totalWords"`
	TotalRequests       int `json:"totalRequests"`
	AvgProcessingTimeNs int `json:"avgProcessingTimeNs"`
}
