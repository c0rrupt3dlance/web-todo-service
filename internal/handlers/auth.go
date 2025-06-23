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

type signInUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) SignIn(c *gin.Context) {
	var inputUser signInUser

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		log.Println("user sent incorrect input data, unable to bind")
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "incorrect data in input",
		})
		return
	}

	token, err := h.services.GenerateToken(inputUser.Username, inputUser.Password)

	if err != nil {
		log.Printf("error from service: %s", err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
