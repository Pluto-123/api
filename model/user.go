package model

import "gorm.io/gorm"

type User struct {
	Id       int    `gorm:"primary_key;auto_increment" json:"id"`
	Username string `gorm:"size:50" json:"username"`
	Password string `gorm:"size:50" json:"password"`
}

func GetUserById(db *gorm.DB, id int) (*User, error) {
	var user = User{Id: id}
	result := db.First(&user)
	return &user, result.Error
}

func GetUserByName(db *gorm.DB, name string) (*User, error) {
	var user User
	result := db.Where(&User{Username: name}).First(&user)
	return &user, result.Error
}
