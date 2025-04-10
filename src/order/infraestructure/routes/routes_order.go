package routes

import (
	usecases "api/src/order/application/use_cases"
	"api/src/order/domain/repositories"
	"api/src/order/infraestructure/adapters"
	"api/src/order/infraestructure/controllers"

	"log"

	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, repo repositories.IOrder) {
	rabbitProducer, err := adapters.NewRabbitProducer("amqp://leo:1234@34.235.202.211:5672/")
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ producer: %s", err)
	}

	createOrderUseCase := usecases.NewCreateOrder(repo, rabbitProducer)
	createOrderController := controllers.NewCreateOrderController(createOrderUseCase)
	deleteOrderUseCase := usecases.NewDeleteOrder(repo)
	deleteOrderController := controllers.NewDeleteOrderController(deleteOrderUseCase)

	rabbitController := controllers.NewRabbitController(rabbitProducer)

	orderGroup := router.Group("/orders")
	{
		orderGroup.POST("", createOrderController.Run)
		orderGroup.GET("", controllers.GetAllOrdersController(repo))
		orderGroup.DELETE("/:id", deleteOrderController.DeleteOrder)
		orderGroup.POST("/publish", rabbitController.PublishMessage)
	}
}
