package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/port"
)

type PriceHandler struct {
	srv port.PriceService
}

func NewPriceHandler(srv port.PriceService) *PriceHandler {
	return &PriceHandler{srv}
}

func (h PriceHandler) CreatePrice(c *gin.Context) {
	var price domain.PriceRequest

	err := c.ShouldBindJSON(&price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	priceResponse, err := h.srv.CreatePrice(&price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": priceResponse,
	})
}

func (h PriceHandler) UpdatePrice(c *gin.Context) {
	var price domain.PriceRequest

	err := c.ShouldBindJSON(&price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	priceResponse, err := h.srv.UpdatePrice(&price, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": priceResponse,
	})
}

func (h PriceHandler) DeletePrice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	err = h.srv.DeletePrice(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.Status(http.StatusNoContent)
}
