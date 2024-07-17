package month

import (
	"backend/app/pkg/validator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Controller(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/v1")

	v1.GET("/months", func(c *gin.Context) {
		GetMonths(db, c)
	})

	v1.POST("/months", func(c *gin.Context) {
		var reqBody MonthCreate
		if err := validator.BindBody(c, &reqBody); err {
			return
		}

		CreateMonth(db, c, &reqBody)
	})

	v1.PUT("/months/:id", func(c *gin.Context) {
		var reqUri MonthGetById
		var reqBody MonthUpdate
		if err := validator.BindBody(c, &reqBody); err {
			return
		}
		if err := validator.BindUri(c, &reqUri); err {
			return
		}

		UpdateShowMonth(db, c, reqUri, &reqBody)
	})

	v1.DELETE("/months/:id", func(c *gin.Context) {
		var reqUri MonthDelete
		if err := validator.BindUri(c, &reqUri); err {
			return
		}

		DeletMonth(db, c, reqUri)
	})
}

func responseFormatter(month *Month) *MonthResponse {
	return &MonthResponse{
		ID:    month.ID,
		Year:  month.Year,
		Month: month.Month,
		Show:  &month.Show,
	}
}
