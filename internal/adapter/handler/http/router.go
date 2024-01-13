package handler

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	roleHander RoleHandler,
	userHandler UserHandler,
	tagHandler TagHandler,
	goodsHandler GoodsHandler,
) (*Router, error) {
	// Disable debug mode and write logs to file in production
	env := os.Getenv("APP_ENV")
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)

		logFile, _ := os.Create("gin.log")
		gin.DefaultWriter = io.Writer(logFile)
	}

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(customLogger), gin.Recovery())
	v1 := router.Group("api/v1")
	{
		user := v1.Group("/users")
		{
			user.GET("/", userHandler.GetUsers)
			user.GET("/:id", userHandler.GetUserById)
		}
		role := v1.Group("/roles")
		{
			role.GET("/", roleHander.GetRoles)
			role.GET("/:id", roleHander.GetRoleById)
		}
		tag := v1.Group("/tags")
		{
			tag.GET("/", tagHandler.GetTags)
			tag.GET("/:id", tagHandler.GetTagById)
			tag.POST("/", tagHandler.CreateTag)
			tag.PUT("/:id", tagHandler.UpdateTag)
			tag.DELETE("/:id", tagHandler.DeleteTag)
		}
		goods := v1.Group("/goodses")
		{
			goods.GET("/", goodsHandler.GetGoodses)
			goods.GET("/:id", goodsHandler.GetGoodsById)
			goods.POST("/", goodsHandler.CreateGoods)
			goods.PUT("/:id", goodsHandler.UpdateGoods)
			goods.DELETE("/:id", goodsHandler.DeleteGoods)
		}
	}

	return &Router{
		router,
	}, nil
}

// customLogger is a custom Gin logger
func customLogger(param gin.LogFormatterParams) string {
	return fmt.Sprintf("[%s] - %s \"%s %s %s %d %s [%s]\"\n",
		param.TimeStamp.Format(time.RFC1123),
		param.ClientIP,
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency.Round(time.Millisecond),
		param.Request.UserAgent(),
	)
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
