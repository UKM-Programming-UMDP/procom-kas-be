package user

import (
	"backend/app/pkg/filter/datefilter"
	"backend/app/pkg/filter/singlesearch"
	"backend/app/pkg/filter/stringfilter"
	"backend/app/pkg/handler"
	"backend/app/pkg/pagination"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userFilters struct {
	Pagination   handler.Pagination
	DateFilter   datefilter.DateFilter
	StringFilter stringfilter.StringFilter
	NameFilter   searchByName
}

type searchByName struct {
	Name string `form:"name"`
}

func filterService(c *gin.Context, query *gorm.DB, filters *userFilters) {
	datefilter.Build(c, query, &filters.DateFilter)
	stringfilter.Build(c, query, "name", &filters.StringFilter)
	singlesearch.Build(c, query, "name", &filters.NameFilter.Name)

	pagination.Build(c, query, &filters.Pagination)
}
