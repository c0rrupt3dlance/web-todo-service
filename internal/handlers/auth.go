package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"web-todo-service/internal/models"
)

func (h *Handler) SignUp(c *gin.Context) {
	var (
		inputUser models.User
		id        int
	)

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		log.Println("user sent incorrect input data, unable to bind")
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "incorrect data in input",
		})
		return
	}

	id, err := h.services.CreateUser(inputUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": id})
}

func (h *Handler) SignIn(c *gin.Context) {
}
