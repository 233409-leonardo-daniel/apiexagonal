package controllers

import (
	"api/src/user/application/use_cases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsersController(useCase *usecases.ViewUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := useCase.Execute()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}
