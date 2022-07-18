package models

type Ticket struct {
	ID			uint 		`json:"id"	gorm:"primary_key; auto_increment; not_null"`
	UserID		uint		`json:"user_id" gorm:"foreignKey:UserID"`
	EventID		uint		`json:"event_id" gorm:"foreignKey:EventID"`
	Price		string		`json:"price"`
}	