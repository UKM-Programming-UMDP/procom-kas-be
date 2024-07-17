package financialrequest

import (
	"backend/app/api/enums"
	"backend/app/api/user"
)

func responseFormatter(finReq *FinancialRequest) *FinancialRequestResponse {
	return &FinancialRequestResponse{
		RequestID: finReq.RequestID,
		Amount:    finReq.Amount,
		Note:      finReq.Note,
		User: user.UserResponse{
			NPM:   finReq.User.NPM,
			Name:  finReq.User.Name,
			Email: finReq.User.Email,
		},
		Payment: struct {
			Status         enums.PaymentStatus `json:"status,omitempty"`
			Type           enums.PaymentType   `json:"type,omitempty"`
			TargetProvider string              `json:"target_provider,omitempty"`
			TargetName     string              `json:"target_name,omitempty"`
			TargetNumber   string              `json:"target_number,omitempty"`
			Evidence       string              `json:"evidence,omitempty"`
		}{
			Status:         finReq.Payment.Status,
			Type:           finReq.Payment.Type,
			TargetProvider: finReq.Payment.TargetProvider,
			TargetName:     finReq.Payment.TargetName,
			TargetNumber:   finReq.Payment.TargetNumber,
			Evidence:       finReq.Payment.Evidence,
		},
		TransferedEvidence: finReq.TransferedEvidence,
		CreatedAt:          finReq.CreatedAt,
	}
}
