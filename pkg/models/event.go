package models

type Event struct {
	ID			uint 		`json:"id"	gorm:"primary_key; auto_increment; not_null"`
	Band_Name	string		`json:"band_name"`
	Location	string		`json:"location"`
	Price		string		`json:"price"`
	Capacity	int			`json:"capacity"`
	Date 		string		`json:"date"`
}	