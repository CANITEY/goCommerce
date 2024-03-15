package models

type Order struct {
	Cutomer    User
	Products   []Product
	TotalPrice float64
}
