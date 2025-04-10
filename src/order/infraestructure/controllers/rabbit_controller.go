package controllers

import (
	"api/src/order/domain/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RabbitController struct {
	rabbitProducer repositories.IRabbitProducer
}

func NewRabbitController(rabbitProducer repositories.IRabbitProducer) *RabbitController {
	return &RabbitController{rabbitProducer: rabbitProducer}
}

func (rc *RabbitController) PublishMessage(c *gin.Context) {
	var input struct {
		IdProduct  int32   `json:"idProduct"`
		Quantity   int32   `json:"quantity"`
		TotalPrice float64 `json:"totalPrice"`
		Status     string  `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := rc.rabbitProducer.Publish(input.IdProduct, input.Quantity, input.TotalPrice, input.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message published successfully"})
}
