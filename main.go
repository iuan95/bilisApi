package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"bilis/config"
	"bilis/database"
	"bilis/models"
	"bilis/repository"
)
func main(){
	err:=godotenv.Load()
	if err != nil {
		log.Fatal("cant load .env")
	}
	var config = config.DataBaseConfig{
		Host : os.Getenv("HOST"),
		User: os.Getenv("USER"),
		DBName : os.Getenv("DBNAME"),
		Port: os.Getenv("PORT"),
	}
	db, err:= database.ConnectToDatabase(config)
	if(err!= nil){
		fmt.Fprintf(os.Stderr, "Error %v\n", err)
		os.Exit(1)
	}
	r := repository.Repository{
		DB: db,
	}
	models.AutoMigrate(r.DB)
	e := echo.New()
	r.InitRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}