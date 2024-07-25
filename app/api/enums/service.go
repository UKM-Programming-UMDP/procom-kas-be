package enums

import (
	"backend/app/common/models"
	"backend/app/common/utils"
	"backend/app/pkg/handler"
	"backend/app/pkg/validator"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetEnums(db *gorm.DB, c *gin.Context, tableName string) {
	var enums []models.Enums
	table := fmt.Sprintf("%s.%s", utils.GetEnv("environment"), tableName)
	query := db.Table(table).Find(&enums)
	val, _ := validator.Query(query, c, false)
	if !val {
		return
	}

	result := make([]*EnumsResponse, len(enums))
	for i, enum := range enums {
		result[i] = enumResponseFormatter(&enum)
	}

	var response interface{} = result
	handler.Success(c, http.StatusAccepted, "Enums retrieved successfully", &response)
}

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
