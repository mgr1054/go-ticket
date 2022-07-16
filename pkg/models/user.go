package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID		 	uint 		`json:"id" gorm:"primary_key; auto_increment; not_null"`
	Name     	string 		`json:"name"`
	Username 	string 		`json:"username" gorm:"unique"`
	Email    	string 		`json:"email" gorm:"unique"`
	Password 	string 		`json:"password"`
}