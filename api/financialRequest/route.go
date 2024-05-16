package financialRequest

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(router *gin.Engine, db *gorm.DB) {
	controller := Controller{
		Db: db,
	}
	r := router.Group("/api/financial-request")
	r.GET("", controller.GetFinancialRequest)
	r.GET("/details", controller.GetFinancialRequestById)
	r.POST("", controller.CreateFinancialRequest)
	r.PUT("/", controller.UpdateFinancialRequestStatus)
}
