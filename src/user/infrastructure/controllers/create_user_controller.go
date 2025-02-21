package controllers

import (
	"api/src/user/application/use_cases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUserController struct {
	useCase *usecases.CreateUser
}

func NewCreateUserController(useCase *usecases.CreateUser) *CreateUserController {
	return &CreateUserController{useCase: useCase}
}

func (cp *CreateUserController) Run(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		LastName string `json:"lastname"`
		Age      int    `json:"age"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cp.useCase.Execute(input.Name, input.LastName, input.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
