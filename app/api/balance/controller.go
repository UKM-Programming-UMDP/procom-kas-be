package balance

import (
	"backend/app/pkg/validator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Controller(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/v1")

	v1.GET("/balance", func(c *gin.Context) {
		GetBalance(db, c)
	})

	v1.PUT("/balance", func(c *gin.Context) {
		var reqBody BalanceUpdate
		if err := validator.BindBody(c, &reqBody); err {
			return
		}

		UpdateBalance(db, c, &reqBody)
	})

	v1.GET("/balance/history", func(c *gin.Context) {
		var filters balanceFilters
		if err := validator.BindParams(c, &filters); err {
			return
		}

		GetHistories(db, c, &filters)
	})

}
