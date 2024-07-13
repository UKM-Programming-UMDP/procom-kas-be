package main

import (
	"backend/app/common/utils"
	"backend/app/config/database"
	"backend/app/config/mailer"
	"backend/app/config/other"
	"backend/app/config/router"
)

func main() {
	utils.LoadEnv(".env")

	other.InitUploadFolder()
	database.InitializeDB()
	router.InitializeRouter()
	router.InitializeRoutes()
	mailer.InitializeMailer()

	routerInstance := router.GetRouterInstance()
	routerInstance.Run(
		utils.GetEnv("BE_HOST") + ":" + utils.GetEnv("BE_PORT"),
	)
}
