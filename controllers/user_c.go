package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lMikadal/POS_Backend_golang.git/models"
)

func (db *DBController) GetAllUsers(c *gin.Context) {
	var users []models.User
	var userResponse []models.UserResponse

	db.Database.Model(&models.User{}).Preload("Role").Find(&users)

	for _, e := range users {
		userResponse = append(userResponse, models.UserResponse{
			UserName:  e.UserName,
			UserEmail: e.UserEmail,
			RoleName:  e.Role[0].RoleName,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"users": userResponse,
	})
}
