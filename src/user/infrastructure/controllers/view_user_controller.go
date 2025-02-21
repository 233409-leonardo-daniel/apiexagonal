package controllers

import (
	"api/src/user/domain/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllPrductsController(repo repositories.IUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := repo.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}
