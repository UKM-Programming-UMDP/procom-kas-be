package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(router *gin.Engine, db *gorm.DB) {
	controller := Controller{
		Db: db,
	}
	r := router.Group("/api/users")
	r.GET("", controller.GetUsers)
	r.GET("/details", controller.GetUserById)
	r.POST("", controller.CreateUser)
	r.PUT("", controller.UpdateUser)
	r.DELETE("", controller.DeleteUser)
}
