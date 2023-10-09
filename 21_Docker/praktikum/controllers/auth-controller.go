package controllers

import (
	"net/http"
	"praktikum/controllers/requests"
	"praktikum/helpers"
	"praktikum/repositories"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	var loginReq = requests.LoginRequest{}
	errBind := c.Bind(&loginReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error bind "+errBind.Error()))
	}

	data, token, err := repositories.CheckLogin(loginReq.Email, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error login"))
	}
	response := map[string]any{
		"token":   token,
		"user_id": data.ID,
		"name":    data.Name,
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success recieve user data", response))
}
