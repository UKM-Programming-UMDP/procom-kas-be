package kassubmission

import (
	"backend/app/api/enums"
	"backend/app/api/user"
	"backend/app/common/models"
	"time"
)

type KasSubmission struct {
	ID           int                 `gorm:"primaryKey;autoIncrement"`
	SubmissionID string              `gorm:"not null;size:5;unique" validate:"len=5"`
	UserID       int                 `gorm:"not null"`
	User         user.User           `gorm:"foreignKey:UserID;references:id"`
	PayedAmount  int                 `gorm:"not null"`
	StatusID     int                 `gorm:"not null"`
	Status       enums.PaymentStatus `gorm:"foreignKey:StatusID;references:id"`
	Note         string              `gorm:"not null"`
	Evidence     string              `gorm:"not null"`
	models.TimestampsSoftDelete
}

type KasSubmissionGetByID struct {
	SubmissionID string `uri:"id" validate:"required,len=5"`
}

type KasSubmissionCreate struct {
	User        user.User `json:"user" validate:"required"`
	PayedAmount *int      `json:"payed_amount" validate:"required,min=1"`
	Note        string    `json:"note" validate:"required"`
	Evidence    string    `json:"evidence" validate:"required"`
}

type KasSubmissionUpdate struct {
	Status enums.PaymentStatus `json:"status" validate:"required"`
}

type KasSubmissionResponse struct {
	SubmissionID string              `json:"submission_id,omitempty"`
	User         user.UserResponse   `json:"user,omitempty"`
	PayedAmount  *int                `json:"payed_amount,omitempty"`
	Status       enums.PaymentStatus `json:"status,omitempty"`
	Note         string              `json:"note,omitempty"`
	Evidence     string              `json:"evidence,omitempty"`
	SubmittedAt  time.Time           `json:"submitted_at,omitempty"`
	UpdatedAt    time.Time           `json:"updated_at,omitempty"`
}
