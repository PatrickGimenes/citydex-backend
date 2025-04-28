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

func (u *UserController) GetUserById(c *gin.Context) {
	id := c.Param("id_user")

	if id == "" {
		response := model.Response{
			Message: "Id do usuário não pode ser nulo",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := u.userUsecase.GetUserById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if user == nil {
		response := model.Response{
			Message: "User not found",
		}
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, user)

}
