package database

import (
	"backend/app/common/utils"
	"log"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) {
	initBalance(db)
	initPaymentStatus(db)
	initPaymentType(db)
}

func initBalance(db *gorm.DB) {
	tableName := utils.GetEnv("environment") + ".balances"

	var isExist bool
	queryIsBalanceExist := db.Table(tableName).Select("count(*) > 0").Where("id = 1").Scan(&isExist)
	if queryIsBalanceExist.Error != nil {
		panic(queryIsBalanceExist.Error)
	}
	if isExist {
		return
	}

	insert := "INSERT INTO " + tableName
	query := db.Exec(insert + " (id, balance, created_at, updated_at, deleted_at) VALUES (1, 0, now(), now(), null)")
	if query.Error != nil {
		panic(query.Error)
	}
	log.Println("Balance initialized")
}

func initPaymentStatus(db *gorm.DB) {
	tableName := utils.GetEnv("environment") + ".payment_statuses"

	var isExist bool
	queryIsPaymentStatusExist := db.Table(tableName).Select("count(*) > 0").Where("id = 1").Scan(&isExist)
	if queryIsPaymentStatusExist.Error != nil {
		panic(queryIsPaymentStatusExist.Error)
	}
	if isExist {
		return
	}

	insert := "INSERT INTO " + tableName
	query := db.Exec(insert + " (id, name) VALUES (1, 'Approved')")
	if query.Error != nil {
		panic(query.Error)
	}
	query = db.Exec(insert + " (id, name) VALUES (2, 'Rejected')")
	if query.Error != nil {
		panic(query.Error)
	}
	query = db.Exec(insert + " (id, name) VALUES (3, 'Pending')")
	if query.Error != nil {
		panic(query.Error)
	}
	log.Println("Payment Status initialized")
}

func initPaymentType(db *gorm.DB) {
	tableName := utils.GetEnv("environment") + ".payment_types"

	var isExist bool
	queryIsPaymentTypeExist := db.Table(tableName).Select("count(*) > 0").Where("id = 1").Scan(&isExist)
	if queryIsPaymentTypeExist.Error != nil {
		panic(queryIsPaymentTypeExist.Error)
	}
	if isExist {
		return
	}

	insert := "INSERT INTO " + tableName
	query := db.Exec(insert + " (id, name) VALUES (1, 'Transfer')")
	if query.Error != nil {
		panic(query.Error)
	}
	query = db.Exec(insert + " (id, name) VALUES (2, 'Cash')")
	if query.Error != nil {
		panic(query.Error)
	}
	query = db.Exec(insert + " (id, name) VALUES (3, 'Other')")
	if query.Error != nil {
		panic(query.Error)
	}
	log.Println("Payment Type initialized")
}
