package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"web-todo-service/internal/models"
)

func (h *Handler) GetAllItems(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "no user id"})
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	items, err := h.services.TodoItem.GetAll(userId.(int), listId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (h *Handler) GetItemById(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "no user id"})
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	item, err := h.services.TodoItem.GetById(userId.(int), listId, itemId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"item": item})
}

func (h *Handler) AddItem(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "no user id"})
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	var item models.TodoItem

	if err = c.ShouldBindJSON(&item); err != nil {
		logrus.Printf("invalid item %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid item"})
		return
	}

	id, err := h.services.TodoItem.Create(userId.(int), listId, item)

	if err != nil {
		logrus.Printf("error from service: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) UpdateItem(c *gin.Context) {}

func (h *Handler) DeleteItem(c *gin.Context) {}
