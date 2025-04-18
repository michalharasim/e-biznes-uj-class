package routes

import (
	"zad4/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	products := e.Group("/products")
	products.POST("", controllers.CreateProduct)
	products.GET("", controllers.GetProducts)
	products.GET("/:id", controllers.GetProduct)
	products.PUT("/:id", controllers.UpdateProduct)
	products.DELETE("/:id", controllers.DeleteProduct)
}