package routes

import (
	"api/src/product/application/use_cases"
	"api/src/product/domain/repositories"
	"api/src/product/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.Engine, repo repositories.IProduct) {
	createProductCaseUse := usecases.NewCreateProduct(repo)
	createProductController := controllers.NewCreateProductController(createProductCaseUse)
	updateProductUseCase := usecases.NewUpdateProduct(repo)
	updateProductController := controllers.NewUpdateProductController(updateProductUseCase)
	deleteProductUseCase := usecases.NewDeleteProduct(repo)
	deleteProductController := controllers.NewDeleteProductController(deleteProductUseCase)

	productGroup := router.Group("/products")

	{
		productGroup.POST("", createProductController.Run)
		productGroup.GET("", controllers.GetAllPrductsController(repo))
		productGroup.PUT("/:id", updateProductController.UpdateProduct)    // Actualizar un producto
		productGroup.DELETE("/:id", deleteProductController.DeleteProduct) // Nueva ruta

	}
}
