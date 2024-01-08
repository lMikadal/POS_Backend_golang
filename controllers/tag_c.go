package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lMikadal/POS_Backend_golang.git/models"
	"github.com/lMikadal/POS_Backend_golang.git/utils"
)

func (db *DBController) TagGetAll(c *gin.Context) {
	var tags []models.Tag

	db.Database.Find(&tags)

	c.JSON(http.StatusOK, gin.H{
		"data": tags,
	})
}

func (db *DBController) TagGetByID(c *gin.Context) {
	var tag models.Tag

	ok := db.Database.First(&tag, c.Param("id")).Error
	if utils.ErrBadRequest(ok, "tag not found", c) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tag,
	})
}

func (db *DBController) TagCreate(c *gin.Context) {
	var jsonTag models.Tag

	ok := c.ShouldBindJSON(&jsonTag)
	if utils.ErrBadRequest(ok, "invalid json", c) {
		return
	}

	db.Database.Save(&jsonTag)

	tag := models.TagResponse{
		TagID:   jsonTag.ID,
		TagName: jsonTag.TagName,
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": tag,
	})
}
