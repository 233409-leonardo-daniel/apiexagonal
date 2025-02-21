package controllers

import (
	"net/http"
	"strconv"

	"api/src/product/application/use_cases"
	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	useCase *usecases.DeleteProduct
}

func NewDeleteProductController(useCase *usecases.DeleteProduct) *DeleteProductController {
	return &DeleteProductController{useCase: useCase}
}

func (dc *DeleteProductController) DeleteProduct(c *gin.Context) {
	// Obtener el ID del producto desde la URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de producto inválido"})
		return
	}

	// Ejecutar el caso de uso
	err = dc.useCase.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado exitosamente"})
}
