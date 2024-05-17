package config_router

import (
	"backend/api/balance"
	"backend/api/balanceHistory"
	"backend/api/fileUpload"
	"backend/api/financialRequest"
	"backend/api/kasSubmission"
	"backend/api/month"
	"backend/api/user"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(corsConfig())

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	return router
}

func InitRoutes(router *gin.Engine, db *gorm.DB) {
	month.Routes(router, db)
	user.Routes(router, db)
	kasSubmission.Routes(router, db)
	fileUpload.Routes(router)
	balance.Routes(router, db)
	balanceHistory.Routes(router, db)
	financialRequest.Routes(router, db)
}

func corsConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	}
}
