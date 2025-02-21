package controllers

import (
	"net/http"
	"strconv"

	"api/src/user/application/use_cases"
	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	useCase *usecases.UpdateUser
}

func NewUpdateUserController(useCase *usecases.UpdateUser) *UpdateUserController {
	return &UpdateUserController{useCase: useCase}
}

func (uc *UpdateUserController) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	// Parsear el cuerpo de la solicitud
	var input struct {
		Name     string `json:"name"`
		LastName string `json:"lastname" gorm:"column:lastname"`
		Age      int    `json:"age"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input data"})
		return
	}

	err = uc.useCase.Execute(int32(id), input.Name, input.LastName, input.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retornar una respuesta exitosa
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
