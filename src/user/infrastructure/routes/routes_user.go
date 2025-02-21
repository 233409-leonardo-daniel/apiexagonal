package routes

import (
	usecases "api/src/user/application/use_cases"
	"api/src/user/domain/repositories"
	"api/src/user/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, repo repositories.IUser) {
	createUserCaseUse := usecases.NewCreateUser(repo)
	createUserController := controllers.NewCreateUserController(createUserCaseUse)
	updateUsertUseCase := usecases.NewUpdateUser(repo)
	updateUserController := controllers.NewUpdateUserController(updateUsertUseCase)
	deleteUserUseCase := usecases.NewDeleteUser(repo)
	deleteUserController := controllers.NewDeleteUserController(deleteUserUseCase)

	userGroup := router.Group("/users")

	{
		userGroup.POST("", createUserController.Run)
		userGroup.GET("", controllers.GetAllPrductsController(repo))
		userGroup.PUT("/:id", updateUserController.UpdateUser)
		userGroup.DELETE("/:id", deleteUserController.DeleteUser)

	}
}
