package controllers

import (
	"net/http"

	"Go-Auth-API/internal/dtos"
	"Go-Auth-API/internal/services"
	"Go-Auth-API/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserServiceInterface
}

func NewUserController(
	userService services.UserServiceInterface,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) FindUserDetailsByID(c *gin.Context) {

	userID, err := utils.GetCurrentUserID(c)
	if err != nil {
		utils.SendUnauthorizedErrResp(c, err.(*utils.UnauthorizedError))
		return
	}

	user, err := uc.userService.FindDetailsByID(uint(userID))
	if err != nil {
		switch e := err.(type) {
		case *utils.ValidationError:
			utils.SendInvalidErrResp(c, e)

		case *utils.ResourceNotFoundError:
			utils.SendNotFoundErrResp(c, e)

		case *utils.UnauthorizedError:
			utils.SendUnauthorizedErrResp(c, e)

		case *utils.ForbiddenError:
			utils.SendForbiddenErrResp(c, e)

		default:
			utils.SendInternalErrResp(c)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User details fetched successfully",
		"user":    user,
	})
}

func (uc *UserController) Register(c *gin.Context) {
	var req dtos.UserFormDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErr := utils.NewValidationError(
			"Invalid request body",
			"INVALID_REQUEST",
			"request",
		)

		utils.SendInvalidErrResp(c, validationErr)
		return
	}

	ID, err := uc.userService.Register(req)
	if err != nil {
		switch e := err.(type) {
		case *utils.ValidationError:
			utils.SendInvalidErrResp(c, e)

		case *utils.ConflictError:
			utils.SendConflictErrResp(c, e)

		case *utils.ResourceNotFoundError:
			utils.SendNotFoundErrResp(c, e)

		case *utils.UnauthorizedError:
			utils.SendUnauthorizedErrResp(c, e)

		case *utils.ForbiddenError:
			utils.SendForbiddenErrResp(c, e)

		default:
			utils.SendInternalErrResp(c)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"ID":      ID,
	})
}
