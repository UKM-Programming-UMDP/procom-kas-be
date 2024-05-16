package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	Page       int `json:"page" form:"page"`
	Limit      int `json:"limit" form:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

func PaginationBuilder(c *gin.Context, query *gorm.DB, count *int64) (Pagination, bool) {
	var pagination Pagination
	if err := BindParamAndValidate(c, &pagination); err {
		return pagination, true
	}

	if pagination.Page > 0 && pagination.Limit > 0 {
		pagination.TotalItems = int(*count)
		pagination.TotalPages = (pagination.TotalItems + pagination.Limit - 1) / pagination.Limit
		query.Limit(pagination.Limit).Offset((pagination.Page - 1) * pagination.Limit)
	}

	return pagination, false
}
