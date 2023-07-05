package app

import (
	"dictionary-service/app/models"
	"github.com/gin-gonic/gin"
)

func (appImpl *App) GetStats(c *gin.Context) {
	stats := appImpl.Stats
	response := models.StatsResponse{
		TotalWords:          stats.CountWords(appImpl.DB.GetData()),
		TotalRequests:       stats.GetTotalRequests(),
		AvgProcessingTimeNs: stats.CalculateAvgReqProcessingTime(),
	}
	c.JSON(200, response)
}
