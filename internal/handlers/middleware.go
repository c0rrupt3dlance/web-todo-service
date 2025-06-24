package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		log.Println("empty authorization header")
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "empty header",
		})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		log.Printf("%s is invalid auth header", header)
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "invalid authorization header",
		})
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		log.Printf("unable to parse token: %s", err)
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "invalid token",
		})
		return
	}

	c.Set(userCtx, userId)
}
