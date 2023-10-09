package controllers

import (
	"praktikum/config"
	"praktikum/helpers"
	"praktikum/models"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all users

func GetUsersController(c echo.Context) error {

	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {

		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error get data"))

	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success get all data", users))

}

// get user by id

func GetUserController(c echo.Context) error {

	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	err := config.DB.First(&user, id).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("User not found"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("User found", user))
}

// create new user

func CreateUserController(c echo.Context) error {

	user := models.User{}

	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {

		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed create user"))

	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success create new user", user))

}

// delete user by id

func DeleteUserController(c echo.Context) error {

	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	err := config.DB.Delete(&user, id).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed delete user"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success delete user", user))

}

// update user by id

func UpdateUserController(c echo.Context) error {

	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))

	user := models.User{}
	err := config.DB.First(&user, id).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed update user"))
	}

	update := new(models.User)
	if err := c.Bind(update); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("User data not valid"))
	}

	user.Name = update.Name
	user.Email = update.Email
	user.Password = update.Password

	config.DB.Save(&user)

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success update user", user))

}
