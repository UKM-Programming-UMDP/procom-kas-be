package datefilter

import (
	"backend/app/common/consts"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DateFilter struct {
	OrderBy consts.OrderBy `json:"order_by" form:"order_by"`
	Sort    consts.Sort    `json:"sort" form:"sort"`
}

func Build(c *gin.Context, query *gorm.DB, filter *DateFilter) {
	if filter.OrderBy != consts.ASC && filter.OrderBy != consts.DESC {
		return
	}

	if filter.Sort != consts.CreatedAt && filter.Sort != consts.UpdatedAt {
		return
	}

	var reversedOrderBy consts.OrderBy
	if filter.OrderBy == consts.ASC {
		reversedOrderBy = consts.DESC
	} else {
		reversedOrderBy = consts.ASC
	}

	query.Order(fmt.Sprintf("%s %s", filter.Sort, reversedOrderBy))
}
