package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrBadRequest(err error, mes string, c *gin.Context) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": mes,
		})
		return true
	}
	return false
}
