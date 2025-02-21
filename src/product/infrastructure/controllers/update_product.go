package controllers

import (
	"net/http"
	"strconv"

	"api/src/product/application/use_cases"
	"github.com/gin-gonic/gin"
)

type UpdateProductController struct {
	useCase *usecases.UpdateProduct
}

func NewUpdateProductController(useCase *usecases.UpdateProduct) *UpdateProductController {
	return &UpdateProductController{useCase: useCase}
}

func (uc *UpdateProductController) UpdateProduct(c *gin.Context) {
	// Obtener el ID del producto desde la URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	// Parsear el cuerpo de la solicitud
	var input struct {
		Name  string  `json:"name"`
		Price float32 `json:"price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input data"})
		return
	}

	// Ejecutar el caso de uso para actualizar el producto
	err = uc.useCase.Execute(int32(id), input.Name, input.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retornar una respuesta exitosa
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}
