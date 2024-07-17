package balance

import (
	"backend/app/api/user"
	"backend/app/common/consts"
	"backend/app/common/models"
	"time"

	"gorm.io/gorm"
)

type Balance struct {
	ID      int `gorm:"primaryKey;autoIncrement" json:"id" validate:"required"`
	Balance int `gorm:"not null"`
	models.Timestamps
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type BalanceHistory struct {
	ID          int             `gorm:"primaryKey;autoIncrement" json:"id" validate:"required"`
	Amount      int             `gorm:"not null"`
	PrevBalance int             `gorm:"not null"`
	Activity    consts.Activity `gorm:"not null"`
	Note        string          `gorm:"not null"`
	UserNPM     string          `gorm:"not null"`
	UserName    string          `gorm:"not null"`
	models.TimestampsSoftDelete
}

type BalanceUpdate struct {
	Amount   *int      `json:"amount" validate:"required,min=1"`
	Activity *int      `json:"activity" validate:"required"`
	Note     string    `json:"note" validate:"required"`
	User     user.User `json:"user" validate:"required"`
}

type BalanceResponse struct {
	Balance   *int      `json:"balance,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type BalanceHistoryAdd struct {
	Amount      *int            `json:"amount" validate:"required"`
	PrevBalance *int            `json:"prev_balance" validate:"required"`
	Activity    consts.Activity `json:"activity" validate:"required"`
	Note        string          `json:"note" validate:"required"`
	User        struct {
		NPM string `json:"npm" validate:"required,len=10"`
	} `json:"user" validate:"required"`
}

type BalanceHistoryResponse struct {
	Amount      int             `json:"amount,omitempty"`
	PrevBalance int             `json:"prev_balance,omitempty"`
	Activity    consts.Activity `json:"activity,omitempty"`
	Note        string          `json:"note,omitempty"`
	User        struct {
		NPM  string `json:"npm,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
