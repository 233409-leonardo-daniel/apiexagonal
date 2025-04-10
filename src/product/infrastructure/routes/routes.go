package routes

import (
	usecases "api/src/product/application/use_cases"
	"api/src/product/domain/repositories"
	"api/src/product/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.Engine, repo repositories.IProduct) {
	// Configura CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Instanciar los casos de uso
	createProductUseCase := usecases.NewCreateProduct(repo)
	createProductController := controllers.NewCreateProductController(createProductUseCase)
	updateProductUseCase := usecases.NewUpdateProduct(repo)
	updateProductController := controllers.NewUpdateProductController(updateProductUseCase)
	deleteProductUseCase := usecases.NewDeleteProduct(repo)
	deleteProductController := controllers.NewDeleteProductController(deleteProductUseCase)
	viewProductUseCase := usecases.NewViewProduct(repo)

	productGroup := router.Group("/products")

	{
		productGroup.POST("", createProductController.Run)
		productGroup.GET("", controllers.GetAllProductsController(viewProductUseCase))
		productGroup.PUT("/:id", updateProductController.UpdateProduct)    // Actualizar un producto
		productGroup.DELETE("/:id", deleteProductController.DeleteProduct) // Nueva ruta
	}
}
