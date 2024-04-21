package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//   HOST='localhost'
//   PORT='5432'
//   USER='admin'
//   PASSWORD='admin'
//   DBNAME='testapi'

var DB *gorm.DB

func ConnectToDatabase() error{
	dsn := "host=localhost user=admin password=admin dbname=bilis port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err !=nil{
		return err
	}
	DB = db
	return nil
}