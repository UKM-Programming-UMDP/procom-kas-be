package database

import (
	"backend/app/api/balance"
	"backend/app/api/enums"
	"backend/app/api/financialrequest"
	"backend/app/api/kassubmission"
	"backend/app/api/month"
	"backend/app/api/user"

	"gorm.io/gorm"
)

type MigratableModel interface {
	Migrate(db *gorm.DB) error
}

var modelList = []interface{}{
	&enums.PaymentStatus{},
	&enums.PaymentType{},

	&user.User{},
	&month.Month{},
	&balance.Balance{},
	&balance.BalanceHistory{},
	&financialrequest.FinancialRequest{},
	&kassubmission.KasSubmission{},
}

func Migrate(db *gorm.DB) error {
	var schemaName string
	err := db.Raw("SELECT schema_name FROM information_schema.schemata WHERE schema_name = 'development'").Scan(&schemaName).Error
	if err != nil {
		panic(err)
	}

	if schemaName == "" {
		if err := db.Exec("CREATE SCHEMA development").Error; err != nil {
			panic(err)
		}
	}

	for _, model := range modelList {
		if err := db.AutoMigrate(model); err != nil {
			panic(err)
		}
	}
	return nil
}
