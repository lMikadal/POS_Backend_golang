package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lMikadal/POS_Backend_golang.git/controllers"
)

func RouterUser(router *gin.RouterGroup, ctrls *controllers.DBController) {
	router.GET("user", ctrls.GetUser)
}
