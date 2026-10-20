
package initializer

import (
	"Go-Auth-API/internal/controllers"
	"Go-Auth-API/internal/repositories"
	"Go-Auth-API/internal/routes"
	"Go-Auth-API/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeApp(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// API version group
	api := router.Group("/api/v1")

	// Initialize repository
	userRepository := repositories.NewUserRepository(db)

	// Initialize service
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService(userRepository)

	// Initialize controller
	userController := controllers.NewUserController(userService)
		authController := controllers.NewAuthController(authService)


	// Register user routes
	routes.UserRoutes(api, userController)
	routes.AuthRoutes(api, authController)

	return router
}
