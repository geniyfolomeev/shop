package handler

import (
	"github.com/gin-contrib/cors"
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
	router.Use(cors.Default())

	api := router.Group("/api")
	{
		api.POST("/register", h.register)
		api.POST("/login", h.login)

		products := api.Group("/products")
		{
			products.GET("/", h.getAllProducts)
			products.GET("/:id", h.getProductById)

			edit := products.Group("/edit", h.userIdentity)
			{
				edit.POST("/", h.createProduct)
				edit.PUT("/:id", h.updateProduct)
				edit.DELETE("/:id", h.deleteProduct)
			}
		}
	}

	return router
}
