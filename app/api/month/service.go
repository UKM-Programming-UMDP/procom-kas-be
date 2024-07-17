package month

import (
	"backend/app/pkg/handler"
	"backend/app/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetMonths(db *gorm.DB, c *gin.Context) {
	var months []Month
	query := db.Model(months).Order("year, month").Find(&months)
	val, _ := validator.Query(query, c, false)
	if !val {
		return
	}

	result := make([]*MonthResponse, len(months))
	for i, month := range months {
		result[i] = responseFormatter(&month)
	}

	handler.Success(c, http.StatusOK, "Success getting a month", result)
}

func CreateMonth(db *gorm.DB, c *gin.Context, reqBody *MonthCreate) {
	var month Month
	queryDuplicateCheck := db.Where("year = ? AND month = ?", reqBody.Year, reqBody.Month).Find(&month)
	val, _ := validator.Query(queryDuplicateCheck, c, false)
	if !val {
		return
	}
	if month.ID != 0 {
		handler.Error(c, http.StatusConflict, "Month already exists for that year")
		return
	}

	month.Year = reqBody.Year
	month.Month = reqBody.Month
	month.Show = true

	queryCreate := db.Create(&month)
	val, _ = validator.Query(queryCreate, c, false)
	if !val {
		return
	}

	handler.Success(c, http.StatusCreated, "Success creating a month", responseFormatter(&month))
}

func UpdateShowMonth(db *gorm.DB, c *gin.Context, reqUri MonthGetById, reqBody *MonthUpdate) {
	var month Month
	queryMonth := db.Where("id = ?", reqUri.ID).Find(&month)
	val, count := validator.Query(queryMonth, c, true)
	if !val {
		return
	}
	if count == 0 {
		handler.Error(c, http.StatusNotFound, "Month not found or registered yet")
		return
	}

	month.Show = *reqBody.Show

	queryUpdate := db.Save(&month)
	val, _ = validator.Query(queryUpdate, c, false)
	if !val {
		return
	}

	handler.Success(c, http.StatusOK, "Success updating show month", responseFormatter(&month))
}

func DeletMonth(db *gorm.DB, c *gin.Context, reqUri MonthDelete) {
	queryDelete := db.Delete(&Month{ID: reqUri.ID})
	val, _ := validator.Query(queryDelete, c, false)
	if !val {
		return
	}
	if queryDelete.RowsAffected == 0 {
		handler.Error(c, http.StatusNotFound, "Month not found or registered yet")
		return
	}

	handler.Success(c, http.StatusNoContent, "", nil)
}
