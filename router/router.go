package router

import (
	"github.com/aabdullahgungor/go-restapi-rabbitmq/controller"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	var productController controller.ProductController

	main := router.Group("api/v1")
	{
		products := main.Group("products")
		{
			products.GET("/", productController.GetAllProducts)
			products.GET("/:id", productController.GetProductById)
			products.POST("/", productController.CreateProduct)
			products.PUT("/", productController.EditProduct)
			products.DELETE("/:id", productController.DeleteProduct)
		}
	}
	return router
}
