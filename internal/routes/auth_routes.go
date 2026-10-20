package routes

import (
	"Go-Auth-API/internal/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(authRouter *gin.RouterGroup,authController *controllers.AuthController) {
	authRouter.POST("/auth/login", authController.Login)
}