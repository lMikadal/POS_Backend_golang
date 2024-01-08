package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lMikadal/POS_Backend_golang.git/controllers"
)

func RouterTag(router *gin.RouterGroup, ctrls *controllers.DBController) {
	router.GET("tag", ctrls.TagGetAll)
	router.GET("tag/:id", ctrls.TagGetByID)
	router.POST("tag", ctrls.TagCreate)
	router.PUT("tag/:id", ctrls.TagUpdate)
}
