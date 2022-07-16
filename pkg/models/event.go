package models

type event struct {
	ID			string 	`json:"id"`
	Band_Name	string	`json:"band_name"`
	Location	string	`json:"location"`
	City		string	`json:"city"`
	Capacity	int		`json:"capacity"`
	Date 		string	`json:"date"`
}