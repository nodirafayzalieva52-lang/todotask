package models

type Users struct {
	ID  uint     `gorm:"primaryKey" json:"id"`
	Name string	`gorm:"size:100" json:"name"`
	Age int     `gorm:"size:100" json:"age"`
}