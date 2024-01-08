package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lMikadal/POS_Backend_golang.git/models"
	"github.com/lMikadal/POS_Backend_golang.git/utils"
)

func (db *DBController) UsersGetAll(c *gin.Context) {
	var users []models.User
	var userResponse []models.UserResponse

	db.Database.Preload("Role").Find(&users)

	for _, e := range users {
		userResponse = append(userResponse, models.UserResponse{
			UserID:    e.ID,
			UserName:  e.UserName,
			UserEmail: e.UserEmail,
			Role: models.RoleResponse{
				RoleID:   e.Role.ID,
				RoleName: e.Role.RoleName,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

func (db *DBController) UserGetOne(c *gin.Context) {
	var user models.User

	ok := db.Database.Preload("Role").First(&user, c.Param("id")).Error
	if utils.ErrBadRequest(ok, "user not found", c) {
		return
	}

	userResponse := models.UserResponse{
		UserID:    user.ID,
		UserName:  user.UserName,
		UserEmail: user.UserEmail,
		Role: models.RoleResponse{
			RoleID:   user.Role.ID,
			RoleName: user.Role.RoleName,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}
