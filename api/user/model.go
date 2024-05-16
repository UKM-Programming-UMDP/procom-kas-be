package user

import (
	"backend/api/month"
	"time"
)

type UserFK struct {
	NPM string `gorm:"unique;not null" json:"npm" validate:"len=10"`
}

type UserSchema struct {
	ID            int                `gorm:"primaryKey;autoIncrement"`
	NPM           string             `gorm:"unique;not null" json:"npm" validate:"len=10"`
	Name          string             `gorm:"not null"`
	Email         string             `gorm:"unique"`
	KasPayed      int                `gorm:"not null"`
	MonthID       int                `gorm:"nullable"`
	MonthStartPay *month.MonthSchema `gorm:"foreignKey:MonthID;references:ID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UserCreate struct {
	NPM           string            `json:"npm" validate:"required,len=10"`
	Name          string            `json:"name" validate:"required,min=3,max=255"`
	Email         string            `json:"email" validate:"required,email"`
	KasPayed      *int              `json:"kas_payed" validate:"required,min=0"`
	MonthStartPay month.MonthSchema `json:"month_start_pay" validate:"required"`
}

type UserUpdate struct {
	Name          string            `json:"name" validate:"required,min=3,max=255"`
	Email         string            `json:"email" validate:"required,email"`
	KasPayed      *int              `json:"kas_payed" validate:"required,min=0"`
	MonthStartPay month.MonthSchema `json:"month_start_pay" validate:"required"`
}

type UserDelete struct {
	ID int
}

type UserResponse struct {
	NPM           string               `json:"npm,omitempty"`
	Name          string               `json:"name,omitempty"`
	Email         string               `json:"email,omitempty"`
	KasPayed      *int                 `json:"kas_payed,omitempty"`
	MonthStartPay *month.MonthResponse `json:"month_start_pay,omitempty"`
}

type UserRequestParam struct {
	NPM string `form:"npm" validate:"required,len=10"`
}
