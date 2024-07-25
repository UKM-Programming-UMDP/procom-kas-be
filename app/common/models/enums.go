package models

type Enums struct {
	ID   int    `gorm:"primaryKey;not null;autoIncrement" validate:"required"`
	Name string `gorm:"not null;size:50"`
}
