package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	service "github.com/lMikadal/POS_Backend_golang.git/services"
)

type roleHandler struct {
	roleService service.RoleService
}

func NewRoleHandler(roleService service.RoleService) roleHandler {
	return roleHandler{roleService: roleService}
}

func (h roleHandler) GetRoles(c *gin.Context) {
	roles, err := h.roleService.GetRoles()
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

func (h roleHandler) GetRoleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	role, err := h.roleService.GetRoleById(id)
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
