package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ventra.com/backend/api/handlers"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	statsHandler := handlers.InitStatsHandler(db)
	leadsHandler := handlers.InitLeadsHandler(db)
	tagsHandler := handlers.InitTagsHandler(db)

	api := router.Group("/api")
	{
		api.GET("/stats", statsHandler.GetStats)
		api.POST("/stats/refresh", statsHandler.RefreshStats)

		api.GET("/leads", leadsHandler.GetLeads)
		api.GET("/leads/:id", leadsHandler.GetLead)
		api.POST("/leads", leadsHandler.CreateLead)
		api.PUT("/leads/:id", leadsHandler.UpdateLead)
		api.DELETE("/leads/:id", leadsHandler.DeleteLead)

		api.GET("/tags", tagsHandler.GetTags)
		api.GET("/tags/:id", tagsHandler.GetTag)
		api.POST("/tags", tagsHandler.CreateTag)
		api.DELETE("/tags/:id", tagsHandler.DeleteTag)
	}

	return router
}