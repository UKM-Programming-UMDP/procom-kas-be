package kasSubmission

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(router *gin.Engine, db *gorm.DB) {
	controller := Controller{
		Db: db,
	}
	r := router.Group("/api/kas")
	r.GET("", controller.GetKasSubmissions)
	r.GET("/details", controller.GetKasSubmissionById)
	r.PUT("", controller.UpdateKasSubmissionStatus)
	r.POST("", controller.CreateKasSubmission)
}
