package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lMikadal/POS_Backend_golang.git/controllers"
)

func SetCollectionRoutes(router *gin.RouterGroup) { //, db *gorm.DB) {
	// ctrls := controllers.DBController{Database: db}

	router.GET("test", controllers.GetTest)
}
