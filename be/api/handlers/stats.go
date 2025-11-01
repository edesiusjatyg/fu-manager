package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ventra.com/backend/models"
)

type StatsHandler struct {
	db *gorm.DB
}

func InitStatsHandler(db *gorm.DB) *StatsHandler {
	return &StatsHandler{db: db}
}

func (handler *StatsHandler) GetStats(ginContext *gin.Context) {
	var stats models.StatsData
	ctx := context.Background()

	result, err := gorm.G[models.StatsData](handler.db).Where("id = ?", 1).First(ctx)
	if err != nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve stats"})
		return
	}

	stats = result

	ginContext.IndentedJSON(http.StatusOK, stats)
}

func (handler *StatsHandler) RefreshStats(ginContext *gin.Context) {
	leadsAlltimeCount := int64(0)
	leadsTodayCount := int64(0)
	followUpTodayCount := int64(0)
	dealsAlltimeCount := int64(0)

	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)
	sevenDaysAgo := startOfDay.Add(-24*7*time.Hour)
	yesterday := startOfDay.Add(-10*time.Minute)
	ctx := context.Background()

	handler.db.Model(&models.LeadsData{}).Count(&leadsAlltimeCount)
	handler.db.Model(&models.LeadsData{}).Where("created_at BETWEEN ? AND ?", startOfDay, endOfDay).Count(&leadsTodayCount)
	handler.db.Model(&models.LeadsData{}).Where("created_at BETWEEN ? AND ?", sevenDaysAgo, yesterday).Count(&followUpTodayCount)

	updated, err := gorm.G[models.StatsData](handler.db).Where("id = ?", 1).Updates(ctx, models.StatsData{
		LeadsAlltime: int(leadsAlltimeCount), 
		LeadsDaily: int(leadsTodayCount), 
		FollowUpToday: int(followUpTodayCount), 
		DealsAlltime:  int(dealsAlltimeCount),
	})

	if err != nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stats"})
		return
	}

	ginContext.IndentedJSON(http.StatusOK, gin.H{"updated": updated})
}