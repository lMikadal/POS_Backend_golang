package routers

import (
	"github.com/gin-gonic/gin"
	handler "github.com/lMikadal/POS_Backend_golang.git/handlers"
	repository "github.com/lMikadal/POS_Backend_golang.git/repositories"
	service "github.com/lMikadal/POS_Backend_golang.git/services"
	"gorm.io/gorm"
)

func RouterUser(db *gorm.DB, router *gin.RouterGroup) {
	userRepository := repository.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router.GET("/users", userHandler.GetUsers)
	router.GET("/user/:id", userHandler.GetUserById)
}
