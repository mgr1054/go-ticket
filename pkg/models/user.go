package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID		 	uint 		`json:"id" gorm:"primary_key; auto_increment; not_null"`
	Name     	string 		`json:"name"`
	Username 	string 		`json:"username" binding:"required" gorm:"unique"`
	Email    	string 		`json:"email" binding:"required" gorm:"unique"`
	Password 	string 		`json:"password binding:"required"`
	Role		string		`json:"role"`
}


func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}