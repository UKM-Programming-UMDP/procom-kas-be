package config_database

import (
	"backend/api/balance"
	"log"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) {
	initBalance(db)
}

func initBalance(db *gorm.DB) {
	var isExist bool
	queryIsBalanceExist := db.Model(balance.BalanceSchema{}).Select("count(*) > 0").Where("id = 1").Find(&isExist)
	if queryIsBalanceExist.Error != nil {
		panic(queryIsBalanceExist.Error)
	}
	if isExist {
		return
	}

	db.Exec("INSERT INTO balance_schemas (id, balance, created_at	, updated_at) VALUES (1, 0, now(), now())")
	log.Println("Balance initialized")
}
