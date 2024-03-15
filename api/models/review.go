package models

type Review struct {
	ID          uint
	Reviewer    User
	Description string
}
