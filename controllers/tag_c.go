package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lMikadal/POS_Backend_golang.git/models"
	"github.com/lMikadal/POS_Backend_golang.git/utils"
)

func (db *DBController) TagGetAll(c *gin.Context) {
	var tags []models.Tag
	var tagResponse []models.TagResponse

	db.Database.Find(&tags)

	for _, e := range tags {
		tagResponse = append(tagResponse, models.TagResponse{
			TagID:   e.ID,
			TagName: e.TagName,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tagResponse,
	})
}

func (db *DBController) TagGetByID(c *gin.Context) {
	var tag models.Tag

	ok := db.Database.First(&tag, c.Param("id")).Error
	if utils.ErrBadRequest(ok, "tag not found", c) {
		return
	}

	tagResponse := models.TagResponse{
		TagID:   tag.ID,
		TagName: tag.TagName,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tagResponse,
	})
}

func (db *DBController) TagCreate(c *gin.Context) {
	var jsonTag models.Tag

	ok := c.ShouldBindJSON(&jsonTag)
	if utils.ErrBadRequest(ok, "invalid json", c) {
		return
	}

	db.Database.Save(&jsonTag)

	tagResponse := models.TagResponse{
		TagID:   jsonTag.ID,
		TagName: jsonTag.TagName,
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": tagResponse,
	})
}
