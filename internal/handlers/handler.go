package handlers

import (
	"github.com/gin-gonic/gin"
	"web-todo-service/internal/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(s *services.Service) *Handler {
	return &Handler{
		services: s,
	}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"ping": "pong",
		})
	})

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}
	api := router.Group("/api/v1", h.userIdentity)
	{
		list := api.Group("/list")
		{
			list.GET("/", h.GetLists)
			list.POST("/", h.AddList)
			list.GET("/:id", h.GetListById)
			list.PUT("/:id", h.UpdateList)
			list.DELETE("/:id", h.DeleteList)
			items := list.Group("/items")
			{
				items.GET("/", h.GetItemsOfList)
				items.POST("/", h.AddItemToList)
				items.GET("/:item_id", h.GetItemById)
				items.PUT("/:item_id", h.UpdateItem)
				items.DELETE("/:item_id", h.DeleteItem)
			}
		}
	}
	return router
}
