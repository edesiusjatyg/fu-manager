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

func InitStatsHandler(dbParam *gorm.DB) *StatsHandler {
	return &StatsHandler{db: dbParam}
}

func (handler *StatsHandler) GetStats(ginContext *gin.Context) {
	var stats models.StatsData
	ctx := context.Background()

	result, err := gorm.G[models.StatsData](handler.db).Where("id = ?", 1).First(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			defaultStats := models.StatsData{
				LeadsAlltime:  0,
				LeadsDaily:    0,
				FollowUpToday: 0,
				DealsAlltime:  0,
			}
			createErr := gorm.G[models.StatsData](handler.db).Create(ctx, &defaultStats)
			if createErr != nil {
				ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize stats"})
				return
			}
			stats = defaultStats
			ginContext.JSON(http.StatusOK, stats)
			return
		}
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve stats"})
		return
	}

	stats = result

	ginContext.JSON(http.StatusOK, stats)
}

func (handler *StatsHandler) RefreshStats(ginContext *gin.Context) {
	leadsAlltimeCount := int64(0)
	leadsTodayCount := int64(0)
	followUpTodayCount := int64(0)
	dealsAlltimeCount := int64(0)

	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)
	sevenDaysAgo := startOfDay.Add(-24 * 7 * time.Hour)
	yesterday := startOfDay.Add(-10 * time.Minute)
	ctx := context.Background()

	handler.db.Model(&models.LeadsData{}).Count(&leadsAlltimeCount)
	handler.db.Model(&models.LeadsData{}).Where("created_at BETWEEN ? AND ?", startOfDay, endOfDay).Count(&leadsTodayCount)
	handler.db.Model(&models.LeadsData{}).Where("created_at BETWEEN ? AND ?", sevenDaysAgo, yesterday).Count(&followUpTodayCount)

	updated, err := gorm.G[models.StatsData](handler.db).Where("id = ?", 1).Updates(ctx, models.StatsData{
		LeadsAlltime:  int(leadsAlltimeCount),
		LeadsDaily:    int(leadsTodayCount),
		FollowUpToday: int(followUpTodayCount),
		DealsAlltime:  int(dealsAlltimeCount),
	})

	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stats"})
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{"updated": updated})
}
