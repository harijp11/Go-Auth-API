package controllers

import (
	"errors"
	"net/http"

	"Go-Auth-API/internal/dtos"
	"Go-Auth-API/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthServiceInterface
}

func NewAuthController(
	authService services.AuthServiceInterface,
) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	var req dtos.LoginFormDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid login request",
		})
		return
	}

	token, err := ac.authService.Login(req)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrInvalidCredentials):
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		case errors.Is(err, services.ErrAuthConfig):
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Authentication configuration error",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Login failed",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}