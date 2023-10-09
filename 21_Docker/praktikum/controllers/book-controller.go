package controllers

import (
	"praktikum/config"
	"praktikum/helpers"
	"praktikum/models"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all books

func GetBooksController(c echo.Context) error {

	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {

		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error get data"))

	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success get all books", books))

}

// get book by id

func GetBookController(c echo.Context) error {

	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	book := models.Book{}
	err := config.DB.First(&book, id).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Book not found"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success get book", book))
}

// create new book

func CreateBookController(c echo.Context) error {

	book := models.Book{}

	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {

		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed create book"))

	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success create new book", book))

}

// delete book by id

func DeleteBookController(c echo.Context) error {

	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	book := models.Book{}
	err := config.DB.Delete(&book, id).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed delete book"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success delete book", book))

}

// update book by id

func UpdateBookController(c echo.Context) error {

	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))

	book := models.Book{}
	err := config.DB.First(&book, id).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed update book"))
	}

	update := new(models.Book)
	if err := c.Bind(update); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Book data is not valid"))
	}

	book.Judul = update.Judul
	book.Penulis = update.Penulis
	book.Penerbit = update.Penerbit

	config.DB.Save(&book)

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success update book", book))

}
