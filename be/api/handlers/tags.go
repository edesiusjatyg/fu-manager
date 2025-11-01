package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ventra.com/backend/models"
)

type TagsHandler struct {
	db *gorm.DB
}

func InitTagsHandler(dbParam *gorm.DB) *TagsHandler {
	return &TagsHandler{db: dbParam}
}

func (handler *TagsHandler) GetTags(ginContext *gin.Context) {
	var tags []models.LeadsTags
	ctx := context.Background()

	result, err := gorm.G[models.LeadsTags](handler.db).Find(ctx)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tags"})
	}

	tags = result

	ginContext.JSON(http.StatusOK, tags)
}

func (handler *TagsHandler) GetTag(ginContext *gin.Context) {
	idParam := ginContext.Param("id")
	var tag models.LeadsTags
	ctx := context.Background()

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag ID"})
		return
	}

	result, err := gorm.G[models.LeadsTags](handler.db).Where("id = ?", id).First(ctx)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tag with id"})
	}

	tag = result

	ginContext.JSON(http.StatusOK, tag)
}

func (handler *TagsHandler) CreateTag(ginContext *gin.Context) {
	var tag models.LeadsTags
	ctx := context.Background()

	if err := ginContext.BindJSON(&tag); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided"})
		return
	}

	err := gorm.G[models.LeadsTags](handler.db).Create(ctx, &tag)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tag"})
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{"message": "Tag created successfully"})
}

func (handler *TagsHandler) DeleteTag(ginContext *gin.Context) {
	idParam := ginContext.Param("id")
	ctx := context.Background()

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag ID"})
		return
	}

	result, err := gorm.G[models.LeadsTags](handler.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tag"})
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{"message": "Tag deleted successfully", "result": result})
}