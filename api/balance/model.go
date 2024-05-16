package balance

import (
	"backend/api/user"
	"time"
)

type BalanceSchema struct {
	ID        int `gorm:"primaryKey;autoIncrement" json:"id" validate:"required"`
	Balance   int `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BalanceUpdate struct {
	Amount   *int            `json:"amount" validate:"required,min=1"`
	Activity *int            `json:"activity" validate:"required"`
	Note     string          `json:"note" validate:"required"`
	User     user.UserSchema `json:"user" validate:"required"`
}

type BalanceResponse struct {
	Balance   *int      `json:"balance,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
