package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ventra.com/backend/api/handlers"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	statsHandler := handlers.InitStatsHandler(db)

	api := router.Group("/api")
	{
		api.GET("/stats", statsHandler.GetStats)
		api.POST("/stats/refresh", statsHandler.RefreshStats)
	}

	return router
}