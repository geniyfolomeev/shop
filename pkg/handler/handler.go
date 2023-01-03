package handler

import (
	"github.com/gin-gonic/gin"
	"shop/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/register", h.register)
		api.POST("/login", h.login)
		products := router.Group("/products")
		{
			products.POST("/", h.createProduct)
			products.GET("/", h.getAllProducts)
			products.GET("/:id", h.getProductById)
			products.PUT("/:id", h.updateProduct)
			products.DELETE("/:id", h.deleteProduct)
		}
	}

	return router
}
