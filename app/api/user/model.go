package user

import (
	"backend/app/api/month"
	"backend/app/common/models"
)

type UserFK struct {
	NPM string `gorm:"unique;not null" json:"npm" validate:"len=10"`
}

type User struct {
	ID            int          `gorm:"primaryKey;autoIncrement"`
	NPM           string       `gorm:"size:10;not null" json:"npm" validate:"len=10"`
	Name          string       `gorm:"not null"`
	Email         string       `gorm:"not null"`
	KasPayed      int          `gorm:"not null"`
	MonthID       int          `gorm:"nullable"`
	MonthStartPay *month.Month `gorm:"foreignKey:MonthID;references:ID"`
	models.TimestampsSoftDelete
}

type UserGetByNPM struct {
	NPM string `uri:"npm" validate:"required,len=10"`
}

type UserCreate struct {
	NPM           string      `json:"npm" validate:"required,len=10"`
	Name          string      `json:"name" validate:"required,min=3,max=255"`
	Email         string      `json:"email" validate:"required,email"`
	KasPayed      *int        `json:"kas_payed" validate:"required,min=0"`
	MonthStartPay month.Month `json:"month_start_pay" validate:"required"`
}

type UserUpdate struct {
	Name          string      `json:"name" validate:"required,min=3,max=255"`
	Email         string      `json:"email" validate:"required,email"`
	KasPayed      *int        `json:"kas_payed" validate:"required,min=0"`
	MonthStartPay month.Month `json:"month_start_pay" validate:"required"`
}

type UserDelete struct {
	NPM string `uri:"npm" validate:"required,len=10"`
}

type UserResponse struct {
	NPM           string               `json:"npm,omitempty"`
	Name          string               `json:"name,omitempty"`
	Email         string               `json:"email,omitempty"`
	KasPayed      *int                 `json:"kas_payed,omitempty"`
	MonthStartPay *month.MonthResponse `json:"month_start_pay,omitempty"`
}
