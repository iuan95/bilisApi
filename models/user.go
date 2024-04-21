package models

import "gorm.io/gorm"


type User struct{
	ID        uint    `gorm:"primary key;autoIncrement" json:"id"`
	Surname    *string `json:"surname"`
	Name     *string `json:"name"`
	Patronymic *string `json:"patronymic"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
	DateOfBirth *string `json:"date_of_birth"`
	FamilyStatus *string `json:"family_status"`
	Country *string `json:"country"`
	City *string `json:"city"`
	District *string `json:"district"`
	Token *string `json:"token"`
	AccessToken *string `json:"access_token"`
	RefreshToken *string `json:"refresh_token"`
}
func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}