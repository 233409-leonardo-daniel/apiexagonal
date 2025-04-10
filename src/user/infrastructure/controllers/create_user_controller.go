package controllers

import (
	usecases "api/src/user/application/use_cases"
	"net/http"

	"github.com/gin-gonic/gin"
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
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cp.useCase.Execute(input.Name, input.LastName, input.Age, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
