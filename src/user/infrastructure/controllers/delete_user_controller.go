package controllers

import (
	"net/http"
	"strconv"

	"api/src/user/application/use_cases"
	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	useCase *usecases.DeleteUser
}

func NewDeleteUserController(useCase *usecases.DeleteUser) *DeleteUserController {
	return &DeleteUserController{useCase: useCase}
}

func (dc *DeleteUserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de Usuario inválido"})
		return
	}

	// Ejecutar el caso de uso
	err = dc.useCase.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usero eliminado exitosamente"})
}
