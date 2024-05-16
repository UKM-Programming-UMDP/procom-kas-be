package balanceHistory

import (
	"time"
)

type Activity string

const (
	Add       Activity = "Add"
	Substract Activity = "Substract"
)

type BalanceHistorySchema struct {
	ID          int      `gorm:"primaryKey;autoIncrement" json:"id" validate:"required"`
	Amount      int      `gorm:"not null"`
	PrevBalance int      `gorm:"not null"`
	Activity    Activity `gorm:"not null"`
	Note        string   `gorm:"not null"`
	UserNPM     string   `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BalanceHistoryAdd struct {
	Amount      *int     `json:"amount" validate:"required"`
	PrevBalance *int     `json:"prev_balance" validate:"required"`
	Activity    Activity `json:"activity" validate:"required"`
	Note        string   `json:"note" validate:"required"`
	User        struct {
		NPM string `json:"npm" validate:"required,len=10"`
	} `json:"user" validate:"required"`
}

type BalanceHistoryResponse struct {
	Amount      int      `json:"amount,omitempty"`
	PrevBalance int      `json:"prev_balance,omitempty"`
	Activity    Activity `json:"activity,omitempty"`
	Note        string   `json:"note,omitempty"`
	User        struct {
		NPM string `json:"npm,omitempty"`
	} `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
