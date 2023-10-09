package routes

import (
	"praktikum/controllers"
	m "praktikum/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	m.LogMiddleware(e)

	// Users
	e.GET("/users", controllers.GetUsersController, m.JWTMiddleware())
	e.POST("/users", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginController)
	e.GET("/users/:id", controllers.GetUserController, m.JWTMiddleware())
	e.DELETE("/users/:id", controllers.DeleteUserController, m.JWTMiddleware())
	e.PUT("/users/:id", controllers.UpdateUserController, m.JWTMiddleware())

	
	// Books
	e.GET("/books", controllers.GetBooksController, m.JWTMiddleware())
	e.POST("/books", controllers.CreateBookController, m.JWTMiddleware())
	e.GET("/books/:id", controllers.GetBookController, m.JWTMiddleware())
	e.DELETE("/books/:id", controllers.DeleteBookController, m.JWTMiddleware())
	e.PUT("/books/:id", controllers.UpdateBookController, m.JWTMiddleware())
	
	return e
}
