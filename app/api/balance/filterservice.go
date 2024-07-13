package balance

import (
	"backend/app/pkg/filter/datefilter"
	"backend/app/pkg/filter/singlesearch"
	"backend/app/pkg/handler"
	"backend/app/pkg/pagination"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type balanceFilters struct {
	Pagination handler.Pagination
	DateFilter datefilter.DateFilter
	NameFilter searchByName
}

type searchByName struct {
	Name string `form:"name"`
}

func filterService(c *gin.Context, query *gorm.DB, filters *balanceFilters) {
	datefilter.Build(c, query, &filters.DateFilter)
	singlesearch.Build(c, query, "user_name", &filters.NameFilter.Name)

	pagination.Build(c, query, &filters.Pagination)
}
