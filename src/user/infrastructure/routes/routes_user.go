package routes

import (
	usecases "api/src/user/application/use_cases"
	"api/src/user/domain/repositories"
	"api/src/user/infrastructure/controllers"
	"api/src/user/infrastructure/services"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, repo repositories.IUser) {
	// Configura CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	bcryptService := services.NewBcryptService()
	createUserCaseUse := usecases.NewCreateUser(repo, bcryptService)
	createUserController := controllers.NewCreateUserController(createUserCaseUse)
	updateUsertUseCase := usecases.NewUpdateUser(repo)
	updateUserController := controllers.NewUpdateUserController(updateUsertUseCase)
	deleteUserUseCase := usecases.NewDeleteUser(repo)
	deleteUserController := controllers.NewDeleteUserController(deleteUserUseCase)
	viewUserUseCase := usecases.NewViewUser(repo)

	userGroup := router.Group("/users")

	{
		userGroup.POST("", createUserController.Run)
		userGroup.GET("", controllers.GetAllUsersController(viewUserUseCase))
		userGroup.PUT("/:id", updateUserController.UpdateUser)
		userGroup.DELETE("/:id", deleteUserController.DeleteUser)
	}
}
