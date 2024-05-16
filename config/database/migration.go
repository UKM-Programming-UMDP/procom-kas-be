package config_database

import (
	"backend/api/balance"
	"backend/api/balanceHistory"
	"backend/api/financialRequest"
	"backend/api/kasSubmission"
	"backend/api/month"
	"backend/api/user"

	"gorm.io/gorm"
)

type MigratableModel interface {
	Migrate(db *gorm.DB) error
}

var modelList = []interface{}{
	&user.UserSchema{},
	&month.MonthSchema{},
	&kasSubmission.KasSubmissionSchema{},
	&balance.BalanceSchema{},
	&balanceHistory.BalanceHistorySchema{},
	&financialRequest.FinancialRequestSchema{},
}

func Migrate(db *gorm.DB) error {
	for _, model := range modelList {
		if err := db.AutoMigrate(model); err != nil {
			panic(err)
		}
	}
	return nil
}
