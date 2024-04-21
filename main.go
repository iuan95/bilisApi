package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)
func main(){
	
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		var Answer struct{
			Message string `json:"message"`
		}
		Answer.Message = "hello ejje"
		return c.JSON(http.StatusOK, &Answer)
	})
	e.Logger.Fatal(e.Start(":1323"))
}