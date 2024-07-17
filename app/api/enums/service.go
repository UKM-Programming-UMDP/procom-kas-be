package enums

import (
	"backend/app/pkg/handler"
	"backend/app/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IsPaymentStatusValid(db *gorm.DB, c *gin.Context, statusID int, isHandleReturn bool) (err bool, val bool) {
	query := db.Model(PaymentStatus{}).Where("id = ?", statusID)
	val, count := validator.Query(query, c, true)
	if !val {
		return true, true
	}
	if count == 0 {
		if isHandleReturn {
			handler.Error(c, http.StatusBadRequest, "Payment status is not valid")
		}
		return false, false
	}

	return false, true
}

func IsPaymentTypeValid(db *gorm.DB, c *gin.Context, typeID int, isHandleReturn bool) (err bool, val bool) {
	query := db.Model(PaymentType{}).Where("id = ?", typeID)
	val, count := validator.Query(query, c, true)
	if !val {
		return true, true
	}
	if count == 0 {
		if isHandleReturn {
			handler.Error(c, http.StatusBadRequest, "Payment type is not valid")
		}
		return false, false
	}

	return false, true
}
