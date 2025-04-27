package controller

import (
	"citydex/model"
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

func (u *UserController) CreateUser(c *gin.Context) {
	var user model.User

	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	insertedUser, err := u.userUsecase.CreateUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, insertedUser)
}
