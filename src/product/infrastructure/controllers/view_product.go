package controllers

import (
	"api/src/product/application/use_cases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProductsController(useCase *usecases.ViewProduct) gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := useCase.Execute()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}
