package models

import (
	"log"

	"gorm.io/gorm"
)



func AutoMigrate(db *gorm.DB){

	err:= MigrateUsers(db)
	if err!= nil{
		log.Fatal("could not migrate db table User")
	}
}