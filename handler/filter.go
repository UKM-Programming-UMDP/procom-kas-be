package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderBy string
type Sort string

const (
	ASC  OrderBy = "asc"
	DESC OrderBy = "desc"
)
const (
	CreatedAt Sort = "created_at"
	UpdatedAt Sort = "updated_at"
	Name      Sort = "name"
)

type Filter struct {
	OrderBy OrderBy `json:"order_by" form:"order_by"`
	Sort    Sort    `json:"sort" form:"sort"`
}

func FilterBuilder(c *gin.Context, query *gorm.DB, fieldName ...string) bool {
	var filter Filter
	if err := BindParamAndValidate(c, &filter); err {
		return true
	}

	isFilter := filter.OrderBy != "" && filter.Sort != ""
	isFilterByDate := filter.Sort == CreatedAt || filter.Sort == UpdatedAt
	isFilterByName := filter.Sort == Name && len(fieldName) > 0

	if isFilter && isFilterByDate {
		query.Order(string(filter.Sort) + " " + string(filter.OrderBy))
	}

	if isFilter && isFilterByName {
		query.Order(fieldName[0] + " " + string(filter.OrderBy))
	}

	return false
}
