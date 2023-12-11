package models

type Group struct {
	ID     uint   `json:"id"`
	Name   string `gorm:"comment:件名" json:"name"`
	Prefix string `gorm:"comment:前缀" json:"prefix"`
	Users  []User
}
