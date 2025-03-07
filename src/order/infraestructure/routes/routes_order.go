package routes

import (
	usecases "api/src/order/application/use_cases"
	"api/src/order/domain/repositories"
	"api/src/order/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, repo repositories.IOrder) {
	createOrderUseCase := usecases.NewCreateOrder(repo)
	createOrderController := controllers.NewCreateOrderController(createOrderUseCase)
	deleteOrderUseCase := usecases.NewDeleteOrder(repo)
	deleteOrderController := controllers.NewDeleteOrderController(deleteOrderUseCase)

	orderGroup := router.Group("/orders")
	{
		orderGroup.POST("", createOrderController.Run)
		orderGroup.GET("", controllers.GetAllOrdersController(repo))
		orderGroup.DELETE("/:id", deleteOrderController.DeleteOrder)
	}
}
