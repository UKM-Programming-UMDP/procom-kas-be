package month

import (
	"backend/app/common/models"
)

type Month struct {
	ID    int  `gorm:"primaryKey;autoIncrement" json:"id" validate:"required"`
	Year  int  `gorm:"not null"`
	Month int  `gorm:"not null"`
	Show  bool `gorm:"not null"`
	models.Timestamps
}

type MonthGetById struct {
	ID int `uri:"id" validate:"required"`
}

type MonthCreate struct {
	Year  int `json:"year" validate:"required,min=2000,max=9999"`
	Month int `json:"month" validate:"required,min=1,max=12"`
}

type MonthUpdate struct {
	Show *bool `json:"show" validate:"required"`
}

type MonthDelete struct {
	ID int `uri:"id" validate:"required"`
}

type MonthResponse struct {
	ID    int   `json:"id,omitempty"`
	Year  int   `json:"year,omitempty"`
	Month int   `json:"month,omitempty"`
	Show  *bool `json:"show,omitempty"`
}
