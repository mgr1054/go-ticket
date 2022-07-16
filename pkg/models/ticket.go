package models

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	ID			uint 		`json:"id"	gorm:"primary_key; auto_increment; not_null"`
	EventID		uint		`json:"event_id"`
	Price		string		`json:"price"`
	Event		Event		`json:"event" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}	