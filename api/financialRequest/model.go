package financialRequest

import (
	"backend/api/user"
	"time"
)

type status string
type paymentType string

const (
	Approved status = "Approved"
	Rejected status = "Rejected"
	Pending  status = "Pending"
)
const (
	Cash     paymentType = "Cash"
	Transfer paymentType = "Transfer"
)

type Payment struct {
	Type           paymentType `gorm:"not null" json:"type"`
	TargetProvider string      `gorm:"not null" json:"target_provider"`
	TargetName     string      `gorm:"not null" json:"target_name"`
	TargetNumber   string      `gorm:"not null" json:"target_number"`
	Evidence       string      `gorm:"not null" json:"evidence"`
}

type FinancialRequestSchema struct {
	ID                 int              `gorm:"primaryKey;autoIncrement"`
	RequestID          string           `gorm:"not null;size:5;unique" validate:"len=5"`
	Amount             int              `gorm:"not null"`
	Note               string           `gorm:"not null"`
	UserNPM            string           `gorm:"not null"`
	User               *user.UserSchema `gorm:"foreignKey:UserNPM;references:npm"`
	Status             status           `gorm:"not null"`
	Payment            Payment          `gorm:"embedded"`
	TransferedEvidence string           `gorm:"not null"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type FinancialRequestCreate struct {
	Amount  int             `json:"amount" validate:"required"`
	Note    string          `json:"note" validate:"required"`
	User    user.UserSchema `json:"user" validate:"required"`
	Payment struct {
		Type           paymentType `json:"type" validate:"required"`
		TargetProvider string      `json:"target_provider" validate:"required"`
		TargetName     string      `json:"target_name" validate:"required"`
		TargetNumber   string      `json:"target_number" validate:"required"`
		Evidence       string      `json:"evidence" validate:"required"`
	} `json:"payment" validate:"required"`
}

type FinancialRequestUpdate struct {
	Status             *int   `json:"status" validate:"required"`
	TransferedEvidence string `json:"transfered_evidence" validate:"required"`
}

type FinancialRequestResponse struct {
	RequestID          string            `json:"request_id"`
	Amount             int               `json:"amount"`
	Note               string            `json:"note"`
	User               user.UserResponse `json:"user,omitempty"`
	Status             status            `json:"status"`
	Payment            Payment           `json:"payment"`
	TransferedEvidence string            `json:"transfered_evidence"`
	CreatedAt          time.Time         `json:"created_at"`
	UpdatedAt          time.Time         `json:"updated_at"`
}

type FinancialRequestParam struct {
	RequestID string `form:"request_id" validate:"required,len=5"`
}
