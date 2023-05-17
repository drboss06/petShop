package handler

import (
	"github.com/gin-gonic/gin"
	"petShop/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		products := api.Group("/products")
		{
			products.POST("/", h.createProduct)
			products.PUT("/:id", h.updateProduct)
			products.GET("/:id", h.getProduct)
			products.DELETE("/:id", h.deleteProduct)
		}
	}

	return router
}
