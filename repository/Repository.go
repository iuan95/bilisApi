package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"bilis/handlers"
)


type Repository struct {
	DB *gorm.DB
}

func (r *Repository) InitRoutes(e *echo.Echo){
	dd:=e.Group("/api")
	dd.POST("/login", handlers.Login)
	dd.POST("/signin", handlers.SignIn)
	dd.POST("/logout", handlers.Logout)
}