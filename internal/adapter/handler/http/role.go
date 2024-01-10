package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/port"
)

type RoleHandler struct {
	srv port.RoleService
}

func NewRoleHandler(srv port.RoleService) *RoleHandler {
	return &RoleHandler{srv}
}

func (h RoleHandler) GetRoles(c *gin.Context) {
	roles, err := h.srv.GetRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": roles,
	})
}

func (h RoleHandler) GetRoleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	role, err := h.srv.GetRoleById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": role,
	})
}
