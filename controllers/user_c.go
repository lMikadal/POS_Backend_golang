package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lMikadal/POS_Backend_golang.git/models"
	"github.com/lMikadal/POS_Backend_golang.git/utils"
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
		"data": userResponse,
	})
}

func (db *DBController) GetUser(c *gin.Context) {
	var user models.User

	ok := db.Database.Model(&models.User{}).Preload("Role").First(&user, c.Param("id")).Error
	if utils.ErrQuery(ok, "user not found", c) {
		return
	}

	userResponse := models.UserResponse{
		UserName:  user.UserName,
		UserEmail: user.UserEmail,
		RoleName:  user.Role[0].RoleName,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}
