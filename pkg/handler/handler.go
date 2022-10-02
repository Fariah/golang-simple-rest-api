package handler

import (
	"github.com/gin-gonic/gin"
	"golang-simple-rest-api/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	//api := router.Group("/api")
	//{
	//	lists := api.Group("/lists")
	//	{
	//		lists.POST("/")
	//		lists.GET("/")
	//		lists.GET("/:id")
	//		lists.PUT("/:id")
	//		lists.DELETE("/:id")
	//
	//		items := lists.Group(":id/items")
	//		{
	//			items.POST("/")
	//			items.GET("/")
	//			items.GET("/:item_id")
	//			items.PUT("/:item_id")
	//			items.DELETE("/:item_id")
	//		}
	//	}
	//}

	return router
}
