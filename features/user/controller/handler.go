package controller

import (
	"belajar-go-echo/features/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase user.UseCaseInterface
}

func New(userUC user.UseCaseInterface) *UserController {
	return &UserController{
		userUsecase: userUC,
	}
}

func (handler *UserController) CreateUser(c echo.Context) error {
	input := new(UserRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := user.UserCore{
		Email:    input.Email,
		Password: input.Password,
	}

	row,errusers:= handler.userUsecase.Insert(data)
	if errusers != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error insert data",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success insert data",
		"data" : row,
	})
}

func (handler *UserController) GetAllUsers(c echo.Context) error {

	data, err := handler.userUsecase.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all data",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "get all data",
		"data":    data,
	})
}

func (handler *UserController) Login(c echo.Context) error {
	input := new(UserRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := user.UserCore{
		Email:    input.Email,
		Password: input.Password,
	}
	
	data,token, err := handler.userUsecase.Login(data.Email,data.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error login",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "login success",
		"email":    data.Email,
		"token": token,
	})
}
