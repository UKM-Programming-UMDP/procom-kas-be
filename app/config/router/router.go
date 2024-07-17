package router

import (
	"backend/app/api/balance"
	"backend/app/api/fileupload"
	"backend/app/api/financialrequest"
	"backend/app/api/kassubmission"
	"backend/app/api/month"
	"backend/app/api/user"
	"backend/app/config/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

var routerInstance *gin.Engine

func InitializeRouter() {
	fmt.Println("===== Initialize Router =====")
	router := gin.Default()
	router.Use(corsConfig())
	router.Use(rateLimiterConfig())

	routerInstance = router
}

func GetRouterInstance() *gin.Engine {
	return routerInstance
}

func InitializeRoutes() {
	fmt.Println("===== Initialize Routes =====")
	router := GetRouterInstance()
	db := database.GetDBInstance()

	user.Controller(router, db)
	month.Controller(router, db)
	balance.Controller(router, db)
	financialrequest.Controller(router, db)
	kassubmission.Controller(router, db)
	fileupload.Controller(router, db)
}
