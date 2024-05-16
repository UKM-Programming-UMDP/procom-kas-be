package kasSubmission

import (
	"backend/api/user"
	"time"
)

type status string

const (
	Pending  status = "Pending"
	Approved status = "Approved"
	Rejected status = "Rejected"
)

type KasSubmissionSchema struct {
	ID           int              `gorm:"primaryKey;autoIncrement"`
	SubmissionID string           `gorm:"not null;size:5;unique" validate:"len=5"`
	UserNPM      string           `gorm:"not null"`
	User         *user.UserSchema `gorm:"foreignKey:UserNPM;references:npm"`
	PayedAmount  int              `gorm:"not null"`
	Status       status           `gorm:"not null"`
	Note         string           `gorm:"nullable"`
	Evidence     string           `gorm:"nullable"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type KasSubmissionCreate struct {
	User        user.UserSchema `json:"user" validate:"required"`
	PayedAmount *int            `json:"payed_amount" validate:"required,min=1"`
	Note        *string         `json:"note" validate:"required"`
	Evidence    string          `json:"evidence" validate:"required"`
}

type KasSubmissionUpdateStatus struct {
	Status *int `form:"status" validate:"required,len=1"`
}

type KasSubmissionResponse struct {
	SubmissionID string            `json:"submission_id,omitempty"`
	User         user.UserResponse `json:"user,omitempty"`
	PayedAmount  *int              `json:"payed_amount,omitempty"`
	Status       status            `json:"status,omitempty"`
	Note         *string           `json:"note,omitempty"`
	Evidence     string            `json:"evidence,omitempty"`
	SubmittedAt  time.Time         `json:"submitted_at,omitempty"`
	UpdatedAt    time.Time         `json:"updated_at,omitempty"`
}

type KasSubmissionRequestParam struct {
	SubmissionID string `form:"submission_id" validate:"required,len=5"`
}
