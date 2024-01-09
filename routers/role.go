package routers

import (
	"github.com/gin-gonic/gin"
	handler "github.com/lMikadal/POS_Backend_golang.git/handlers"
	repository "github.com/lMikadal/POS_Backend_golang.git/repositories"
	service "github.com/lMikadal/POS_Backend_golang.git/services"
	"gorm.io/gorm"
)

func RouterRole(db *gorm.DB, router *gin.RouterGroup) {
	roleRepository := repository.NewRoleRepositoryDB(db)
	roleService := service.NewRoleService(roleRepository)
	roleHandler := handler.NewRoleHandler(roleService)

	router.GET("/roles", roleHandler.GetRoles)
	router.GET("/role/:id", roleHandler.GetRoleById)
}
