package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (db *DBController) GetTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test for api",
	})
}
