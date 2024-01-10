package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/port"
)

type UserHandler struct {
	srv port.UserService
}

func NewUserHandler(srv port.UserService) *UserHandler {
	return &UserHandler{srv}
}

func (h UserHandler) GetUsers(c *gin.Context) {
	users, err := h.srv.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (h UserHandler) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	user, err := h.srv.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
