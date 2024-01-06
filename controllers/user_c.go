package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lMikadal/POS_Backend_golang.git/models"
)

func (db *DBController) GetUser(c *gin.Context) {
	var users []models.User

	db.Database.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
