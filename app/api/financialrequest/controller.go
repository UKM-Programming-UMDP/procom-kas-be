package financialrequest

import (
	"backend/app/pkg/validator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Controller(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/v1")

	v1.GET("/financial-requests", func(c *gin.Context) {
		var filters finreqFilters
		if err := validator.BindParams(c, &filters); err {
			return
		}

		GetFinancialRequests(db, c, &filters)
	})

	v1.GET("/financial-requests/:id", func(c *gin.Context) {
		var reqUri FinancialRequestGetByID
		if err := validator.BindUri(c, &reqUri); err {
			return
		}

		GetFinancialRequestByID(db, c, reqUri)
	})

	v1.POST("/financial-requests", func(c *gin.Context) {
		var reqBody FinancialRequestCreate
		if err := validator.BindBody(c, &reqBody); err {
			return
		}

		CreateFinancialRequest(db, c, &reqBody)
	})

	v1.PUT("/financial-requests/:id", func(c *gin.Context) {
		var reqUri FinancialRequestGetByID
		if err := validator.BindUri(c, &reqUri); err {
			return
		}

		var reqBody FinancialRequestUpdate
		if err := validator.BindBody(c, &reqBody); err {
			return
		}

		UpdateFinancialRequest(db, c, reqUri, &reqBody)
	})
}
