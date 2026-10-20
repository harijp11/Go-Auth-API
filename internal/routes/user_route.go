package routes

import (
	"Go-Auth-API/internal/controllers"
	"Go-Auth-API/internal/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(userRouter *gin.RouterGroup, userController *controllers.UserController) {
	userRouter.POST("/users/register", userController.Register)
	userRouter.GET("/users/profile",middleware.AuthMiddleware(),// userController.Profile,
)
}
