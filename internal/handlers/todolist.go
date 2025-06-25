package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"web-todo-service/internal/models"
)

func (h *Handler) GetAll(c *gin.Context) {
	id, ok := c.Get(userCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "no user id"})
		return
	}

	lists, err := h.services.GetAll(id.(int))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't get lists due to internal errors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"lists": lists})
}

func (h *Handler) GetListById(c *gin.Context) {
	id, ok := c.Get(userCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "no user id"})
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't convert id to int"})
		return
	}

	list, err := h.services.GetById(id.(int), listId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't get list due to internal errors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"list": list})
}

func (h *Handler) AddList(c *gin.Context) {
	id, ok := c.Get(userCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "no user id"})
		return
	}

	var inputList models.TodoList

	if err := c.ShouldBindJSON(&inputList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input data"})
	}

	listId, err := h.services.TodoList.Create(id.(int), inputList)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong during creating list"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"list_id": listId})
}

func (h *Handler) UpdateList(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	var inputList models.TodoList

	if err := c.ShouldBindJSON(&inputList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input data"})
		return
	}

	err = h.services.TodoList.Update(listId, inputList)

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"list": listId})
}

func (h *Handler) DeleteList(c *gin.Context) {}
