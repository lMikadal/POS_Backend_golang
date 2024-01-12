package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/port"
)

type TagHandler struct {
	srv port.TagService
}

func NewTagHandler(srv port.TagService) *TagHandler {
	return &TagHandler{srv}
}

func (h TagHandler) GetTags(c *gin.Context) {
	tags, err := h.srv.GetTags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tags,
	})
}

func (h TagHandler) GetTagById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	tag, err := h.srv.GetTagById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tag,
	})
}

func (h TagHandler) CreateTag(c *gin.Context) {
	var tagRequest domain.TagRequest

	if err := c.ShouldBindJSON(&tagRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	tagResponse, err := h.srv.CreateTag(&tagRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": tagResponse,
	})
}

func (h TagHandler) UpdateTag(c *gin.Context) {
	var tagRequest domain.TagRequest

	if err := c.ShouldBindJSON(&tagRequest); err != nil {
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

	tagResponse, err := h.srv.UpdateTag(&tagRequest, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tagResponse,
	})
}

func (h TagHandler) DeleteTag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = h.srv.DeleteTag(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.Status(http.StatusNoContent)
}
