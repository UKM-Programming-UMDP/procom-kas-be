package month

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(router *gin.Engine, db *gorm.DB) {
	controller := Controller{
		Db: db,
	}
	r := router.Group("/api/month")
	r.GET("", controller.GetMonths)
	r.POST("", controller.CreateMonth)
	r.PUT("", controller.UpdateShowMonth)
	r.DELETE("", controller.DeleteMonth)
}
