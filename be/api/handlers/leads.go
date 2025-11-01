package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ventra.com/backend/models"
)

type LeadsHandler struct {
	db *gorm.DB
}

func InitLeadsHandler(dbParam *gorm.DB) *LeadsHandler {
	return &LeadsHandler{db: dbParam}
}

func (handler *LeadsHandler) GetLeads(ginContext *gin.Context) {
	var leads []models.LeadsData
	ctx := context.Background()

	result, err := gorm.G[models.LeadsData](handler.db).Find(ctx)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve leads"})
	}

	leads = result

	ginContext.JSON(http.StatusOK, leads)
}

func (handler *LeadsHandler) GetLead(ginContext *gin.Context) {
	idParam := ginContext.Param("id")
	var lead models.LeadsData
	ctx := context.Background()

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lead ID"})
		return
	}

	result, err := gorm.G[models.LeadsData](handler.db).Where("id = ?", id).First(ctx)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve lead with id"})
	}

	lead = result

	ginContext.JSON(http.StatusOK, lead)
}

func (handler *LeadsHandler) CreateLead(ginContext *gin.Context) {
	var lead models.LeadsData
	ctx := context.Background()

	if err := ginContext.BindJSON(&lead); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := gorm.G[models.LeadsData](handler.db).Create(ctx, &lead)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new lead"})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"message": "Lead created successfully"})
}

func (handler *LeadsHandler) UpdateLead(ginContext *gin.Context) {
	idParam := ginContext.Param("id")
	var lead models.LeadsData
	ctx := context.Background()

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lead ID"})
		return
	}

	if err := ginContext.BindJSON(&lead); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result, err := gorm.G[models.LeadsData](handler.db).Where("id = ?", id).Updates(ctx, models.LeadsData{
		Name:    lead.Name,
		Company: lead.Company,
		Notes:   lead.Notes,
		Tags:    lead.Tags,
	})
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update lead"})
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{"message": "Lead updated successfully", "result": result})
}

func (handler *LeadsHandler) DeleteLead(ginContext *gin.Context) {
	idParam := ginContext.Param("id")
	ctx := context.Background()

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lead ID"})
		return
	}

	result, err := gorm.G[models.LeadsData](handler.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete lead"})
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{"message": "Lead deleted successfully", "result": result})
}
