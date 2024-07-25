package enums

import "backend/app/common/models"

type PaymentStatus struct {
	models.Enums
}

type PaymentType struct {
	models.Enums
}

type EnumsResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
