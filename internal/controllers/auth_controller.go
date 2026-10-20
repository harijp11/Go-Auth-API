package controllers

import (
	"net/http"

	"Go-Auth-API/internal/dtos"
	"Go-Auth-API/internal/services"
	"Go-Auth-API/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthServiceInterface
}

func NewAuthController(authService services.AuthServiceInterface) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	var req dtos.LoginFormDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErr := utils.NewValidationError(
			"Invalid login request",
			"INVALID_REQUEST",
			"request",
		)

		utils.SendInvalidErrResp(c, validationErr)
		return
	}

	token, err := ac.authService.Login(req)
	if err != nil {
		switch e := err.(type) {
		case *utils.ValidationError:
			utils.SendInvalidErrResp(c, e)

		case *utils.UnauthorizedError:
			utils.SendUnauthorizedErrResp(c, e)

		case *utils.ConflictError:
			utils.SendConflictErrResp(c, e)

		case *utils.ResourceNotFoundError:
			utils.SendNotFoundErrResp(c, e)

		case *utils.ForbiddenError:
			utils.SendForbiddenErrResp(c, e)

		default:
			utils.SendInternalErrResp(c)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}