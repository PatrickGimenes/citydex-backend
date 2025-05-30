package main

import (
	"citydex/controller"
	"citydex/db"
	"citydex/repository"
	"citydex/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Panic(err)
	}

	//Camada de repository
	UserRepository := repository.NewUserRepository(dbConnection)
	//Camada useCase
	UserUsecase := usecase.NewUserUsecase(UserRepository)
	//Camada de controllers
	UserController := controller.NewUserController(UserUsecase)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/users", UserController.GetUsers)
	router.GET("/user/:id_user", UserController.GetUserById)
	router.POST("/register", UserController.CreateUser)
	router.PUT("/update/:id_user", UserController.UpdateUser)

	router.Run(":8080")
}
