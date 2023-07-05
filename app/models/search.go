package models

type SearchQuery struct {
	Word string `form:"word"`
}

type SearchResponse struct {
	Similar []string `json:"similar"`
}
