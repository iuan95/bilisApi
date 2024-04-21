package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error{
	var Answer struct{
		Message string `json:"message"`
	}
	Answer.Message = "login"
	return c.JSON(http.StatusOK, &Answer)
}

func SignIn(c echo.Context) error{
	var Answer struct{
		Message string `json:"message"`
	}
	Answer.Message = "SignIn"
	return c.JSON(http.StatusOK, &Answer)
}
func Logout(c echo.Context) error{
	var Answer struct{
		Message string `json:"message"`
	}
	Answer.Message = "logout"
	return c.JSON(http.StatusOK, &Answer)
}