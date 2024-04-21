package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"bilis/config"
)


func ConnectToDatabase(config config.DataBaseConfig) (*gorm.DB, error){
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err !=nil{
		return  db,err
	}
	return db,nil
}