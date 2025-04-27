package controller

import (
	"citydex/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) UserController {
	return UserController{
		userUsecase: usecase,
	}
}

func (u *UserController) GetUsers(c *gin.Context) {
	users, err := u.userUsecase.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, users)

}
