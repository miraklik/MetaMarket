package db

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"size:255;not null;unique" json:"username"`
	Password  string `gorm:"size:255;not null;" json:"-"`
	Groceries []Grocery
}

func GetUserById(uid uint) (User, error) {
	var user User

	db, err := Setup()

	if err != nil {
		log.Println(err)
		return User{}, err
	}
	if err := db.Preload("Groceries").Where("id=?", uid).Find(&user).Error; err != nil {
		return user, errors.New("user not found")

	}

	return user, nil
}
