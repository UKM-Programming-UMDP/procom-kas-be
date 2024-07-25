package models

type Enums struct {
	ID   int    `gorm:"primaryKey;not null;autoIncrement" json:"id"`
	Name string `gorm:"not null;size:50" json:"name"`
}
