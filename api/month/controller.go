package month

import (
	"backend/handler"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Db *gorm.DB
}

func (a *Controller) GetMonths(c *gin.Context) {
	var months []MonthSchema
	query := a.Db.Model(months).Order("year, month").Find(&months)
	val, _ := handler.QueryValidator(query, c, false)
	if !val {
		return
	}

	result := make([]MonthResponse, len(months))
	for i, month := range months {
		result[i] = monthResFormatter(month)
	}

	handler.Success(c, http.StatusOK, "Success getting a month", result)
}

func (a *Controller) CreateMonth(c *gin.Context) {
	var reqBody MonthCreate
	if err := handler.BindAndValidate(c, &reqBody); err {
		return
	}

	var month MonthSchema
	queryDuplicateCheck := a.Db.Where("year = ? AND month = ?", reqBody.Year, reqBody.Month).Find(&month)
	val, _ := handler.QueryValidator(queryDuplicateCheck, c, false)
	if !val {
		return
	}

	if month.ID != 0 {
		handler.Error(c, http.StatusConflict, "Month already exists for that year")
		return
	}

	month.Year = reqBody.Year
	month.Month = reqBody.Month
	month.Show = false

	queryCreate := a.Db.Create(&month)
	val, _ = handler.QueryValidator(queryCreate, c, false)
	if !val {
		return
	}

	handler.Success(c, http.StatusCreated, "Success creating a month", monthResFormatter(month))
}

func (a *Controller) UpdateShowMonth(c *gin.Context) {
	var reqParam MonthRequestParam
	if err := handler.BindParamAndValidate(c, &reqParam); err {
		return
	}

	var reqBody MonthUpdate
	if err := handler.BindAndValidate(c, &reqBody); err {
		return
	}

	var month MonthSchema
	queryMonth := a.Db.Where("year = ? AND month = ?", reqParam.Year, reqParam.Month).Find(&month)
	val, _ := handler.QueryValidator(queryMonth, c, false)
	if !val {
		return
	}
	if month.ID == 0 {
		handler.Error(c, http.StatusNotFound, "Month not found or registered yet")
		return
	}

	month.Show = *reqBody.Show

	queryUpdate := a.Db.Save(&month)
	val, _ = handler.QueryValidator(queryUpdate, c, false)
	if !val {
		return
	}

	handler.Success(c, http.StatusOK, "Success updating show month", monthResFormatter(month))
}

func (a *Controller) DeleteMonth(c *gin.Context) {
	var reqParam MonthRequestParam
	if err := handler.BindParamAndValidate(c, &reqParam); err {
		return
	}

	var month MonthDelete
	queryDelete := a.Db.Model(MonthSchema{}).Clauses(utils.Returning("id")).Where("year = ? AND month = ?", reqParam.Year, reqParam.Month).Delete(&month)
	val, _ := handler.QueryValidator(queryDelete, c, false)
	if !val {
		return
	}
	if month.ID == 0 {
		handler.Error(c, http.StatusNotFound, "Month not found or registered yet")
		return
	}

	handler.Success(c, http.StatusOK, "Success deleting a month", nil)
}

func monthResFormatter(month MonthSchema) MonthResponse {
	return MonthResponse{
		ID:    month.ID,
		Year:  month.Year,
		Month: month.Month,
		Show:  &month.Show,
	}
}
