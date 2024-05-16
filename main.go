package main

import (
	config_database "backend/config/database"
	config_router "backend/config/router"
)

func main() {
	db := config_database.InitDB()
	router := config_router.InitRouter()
	config_router.InitRoutes(router, db)

	router.Run("localhost:5000")
}
