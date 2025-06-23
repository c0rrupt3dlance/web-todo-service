package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"web-todo-service/internal/models"
)

func (h *Handler) SignUp(c *gin.Context) {
	var inputUser models.User

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		log.Printf("got this error as response: %s", err)
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}

	id, err := h.services.CreateUser(inputUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": id,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
}
