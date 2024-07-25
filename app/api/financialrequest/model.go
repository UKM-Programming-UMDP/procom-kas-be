package financialrequest

import (
	"backend/app/api/enums"
	"backend/app/api/user"
	"backend/app/common/models"
	"time"
)

type Payment struct {
	StatusID       int                 `gorm:"not null"`
	Status         enums.PaymentStatus `gorm:"foreignKey:StatusID;references:ID"`
	TypeID         int                 `gorm:"not null"`
	Type           enums.PaymentType   `gorm:"foreignKey:TypeID;references:ID"`
	TargetProvider string              `gorm:"not null"`
	TargetName     string              `gorm:"not null"`
	TargetNumber   string              `gorm:"not null"`
	Evidence       string              `gorm:"not null"`
}

type FinancialRequest struct {
	ID                 int       `gorm:"primaryKey;autoIncrement"`
	RequestID          string    `gorm:"not null;size:5;unique" validate:"len=5"`
	Amount             int       `gorm:"not null"`
	Note               string    `gorm:"not null"`
	UserID             int       `gorm:"not null"`
	User               user.User `gorm:"foreignKey:UserID;references:ID"`
	Payment            Payment   `gorm:"embedded"`
	TransferedEvidence string    `gorm:"not null"`
	models.TimestampsSoftDelete
}

type FinancialRequestGetByID struct {
	RequestID string `uri:"id" validate:"required,len=5"`
}

type FinancialRequestCreate struct {
	Amount  int       `json:"amount" validate:"required"`
	Note    string    `json:"note" validate:"required"`
	User    user.User `json:"user" validate:"required"`
	Payment struct {
		Type           enums.PaymentType `json:"type" validate:"required"`
		TargetProvider string            `json:"target_provider" validate:"required"`
		TargetName     string            `json:"target_name" validate:"required"`
		TargetNumber   string            `json:"target_number" validate:"required"`
		Evidence       string            `json:"evidence" validate:"required"`
	} `json:"payment" validate:"required"`
}

type FinancialRequestUpdate struct {
	Status             enums.PaymentStatus `json:"status" validate:"required"`
	TransferedEvidence string              `json:"transfered_evidence" validate:"required"`
}

type FinancialRequestResponse struct {
	RequestID string            `json:"request_id,omitempty"`
	Amount    int               `json:"amount,omitempty"`
	Note      string            `json:"note,omitempty"`
	User      user.UserResponse `json:"user,omitempty"`
	Payment   struct {
		Status         enums.PaymentStatus `json:"status,omitempty"`
		Type           enums.PaymentType   `json:"type,omitempty"`
		TargetProvider string              `json:"target_provider,omitempty"`
		TargetName     string              `json:"target_name,omitempty"`
		TargetNumber   string              `json:"target_number,omitempty"`
		Evidence       string              `json:"evidence,omitempty"`
	} `json:"payment,omitempty"`
	TransferedEvidence string    `json:"transfered_evidence"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
}
