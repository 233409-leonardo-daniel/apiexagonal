package controllers

import (
	"net/http"
	"strconv"

	usecases "api/src/order/application/use_cases"

	"github.com/gin-gonic/gin"
)

type DeleteOrderController struct {
	useCase *usecases.DeleteOrder
}

func NewDeleteOrderController(useCase *usecases.DeleteOrder) *DeleteOrderController {
	return &DeleteOrderController{useCase: useCase}
}

func (dc *DeleteOrderController) DeleteOrder(c *gin.Context) {
	// Obtener el ID de la orden desde la URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de orden inválido"})
		return
	}

	// Ejecutar el caso de uso
	err = dc.useCase.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Orden eliminada exitosamente"})
}
