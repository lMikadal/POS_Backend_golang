package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/port"
)

type GoodsHandler struct {
	srv port.GoodsService
}

func NewGoodsHandler(srv port.GoodsService) *GoodsHandler {
	return &GoodsHandler{srv}
}

func (h GoodsHandler) GetGoodses(c *gin.Context) {
	goodses, err := h.srv.GetGoodses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": goodses,
	})
}

func (h GoodsHandler) GetGoodsById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	goods, err := h.srv.GetGoodsById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": goods,
	})
}

func (h GoodsHandler) CreateGoods(c *gin.Context) {
	var goodsRequest domain.GoodsRequest

	if err := c.ShouldBindJSON(&goodsRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	goods, err := h.srv.CreateGoods(&goodsRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": goods,
	})
}

func (h GoodsHandler) UpdateGoods(c *gin.Context) {
	var goodsRequest domain.GoodsRequest

	if err := c.ShouldBindJSON(&goodsRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	goods, err := h.srv.UpdateGoods(&goodsRequest, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": goods,
	})
}

func (h GoodsHandler) DeleteGoods(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	err = h.srv.DeleteGoods(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.Status(http.StatusNoContent)
}
