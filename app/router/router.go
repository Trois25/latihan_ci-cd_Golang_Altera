package router

import (
	m "belajar-go-echo/app/middlewares"
	"belajar-go-echo/features/user/controller"
	"belajar-go-echo/features/user/repository"
	"belajar-go-echo/features/user/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userRepository := repository.New(db)       //menghubungkan data repo ke db
	userUsecase := usecase.New(userRepository) //data pada usecare berdaarkan repository
	userController := controller.New(userUsecase)

	e.GET("/users", userController.GetAllUsers, m.JWTMiddleware())
	e.POST("/users", userController.CreateUser)
	e.POST("/users/login", userController.Login)
}
