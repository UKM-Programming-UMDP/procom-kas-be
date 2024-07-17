package financialrequest

import (
	"backend/app/api/enums"
	"backend/app/pkg/filter/datefilter"
	"backend/app/pkg/handler"
	"backend/app/pkg/pagination"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type finreqFilters struct {
	Pagination    handler.Pagination
	DateFilter    datefilter.DateFilter
	PaymentStatus filterByStatus
}

type filterByStatus struct {
	PaymentStatus int `form:"status"`
}

func filterService(db *gorm.DB, c *gin.Context, query *gorm.DB, filters *finreqFilters) {
	datefilter.Build(c, query, &filters.DateFilter)
	filterByStatusService(db, c, query, filters.PaymentStatus.PaymentStatus)
	pagination.Build(c, query, &filters.Pagination)
}

func filterByStatusService(db *gorm.DB, c *gin.Context, query *gorm.DB, statusID int) {
	if statusID != 0 {
		if err, val := enums.IsPaymentStatusValid(db, c, statusID, true); err || !val {
			return
		}

		query.Where("status_id = ?", statusID)
	}
}
