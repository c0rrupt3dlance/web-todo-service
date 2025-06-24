package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetLists(c *gin.Context) {}

func (h *Handler) GetListById(c *gin.Context) {}

func (h *Handler) AddList(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": id,
	})
}

func (h *Handler) UpdateList(c *gin.Context) {}

func (h *Handler) DeleteList(c *gin.Context) {}
