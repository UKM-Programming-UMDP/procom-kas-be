package enums

import (
	"backend/app/common/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Controller(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/v1")

	v1.GET("/enums/payment-status", func(c *gin.Context) {
		GetEnums(db, c, "payment_statuses")
	})

	v1.GET("/enums/payment-type", func(c *gin.Context) {
		GetEnums(db, c, "payment_types")
	})
}

func enumResponseFormatter(enum *models.Enums) *EnumsResponse {
	return &EnumsResponse{
		ID:   enum.ID,
		Name: enum.Name,
	}
}
