package balance

import (
	"backend/app/api/user"
	"backend/app/common/consts"
	"backend/app/pkg/handler"
	"backend/app/pkg/validator"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBalance(db *gorm.DB, c *gin.Context) {
	balance := []Balance{}

	query := db.Find(&balance)
	if val, _ := validator.Query(query, c, false); !val {
		return
	}
	if len(balance) == 0 {
		handler.Error(c, http.StatusInternalServerError, "Balance not initialized yet")
		return
	}

	handler.Success(c, http.StatusOK, "Success getting a balance", balanceResFormatter(&balance[0]))
}

func GetHistories(db *gorm.DB, c *gin.Context, filters *balanceFilters) {
	var histories []BalanceHistory
	query := db.Model(histories)
	if val, _ := validator.Query(query, c, false); !val {
		return
	}

	filterService(c, query, filters)
	query.Find(&histories)
	result := make([]*BalanceHistoryResponse, len(histories))
	for i, history := range histories {
		result[i] = historyResFormatter(&history)
	}

	var response interface{} = result
	handler.Success(c, http.StatusOK, "Success getting histories", &response, filters.Pagination)
}

func UpdateBalance(db *gorm.DB, c *gin.Context, reqBody *BalanceUpdate) {
	var balance Balance
	result := db.Where("id = 1").Find(&balance)
	if val, _ := validator.Query(result, c, false); !val {
		return
	}

	var user user.User
	queryUserCheck := db.Model(user).Where("npm = ?", reqBody.User.NPM)
	val, count := validator.Query(queryUserCheck, c, true)
	if !val {
		return
	}
	if count == 0 {
		handler.Error(c, http.StatusBadRequest, "User not found")
		return
	}
	queryUserCheck.First(&user)

	activity, err := activityParser(c, reqBody.Activity)
	if err {
		return
	}
	switch activity {
	case consts.Add:
		balance.Balance += *reqBody.Amount
	case consts.Subtract:
		balance.Balance -= *reqBody.Amount
	default:
		handler.Error(c, http.StatusBadRequest, "Invalid activity")
		return
	}

	historyContext := HistoryContext{
		Context:     c,
		Db:          db,
		Amount:      *reqBody.Amount,
		PrevBalance: balance.Balance,
		Note:        reqBody.Note,
		UserNPM:     user.NPM,
		UserName:    user.Name,
		Activity:    activity,
	}
	if err := AddHistory(historyContext); err {
		return
	}

	queryUpdateBalance := db.Save(&balance)
	if val, _ := validator.Query(queryUpdateBalance, c, false); !val {
		return
	}

	handler.Success(c, http.StatusOK, "Success updating balance", balanceResFormatter(&balance))
}

func AddHistory(historyContext HistoryContext) bool {
	history := BalanceHistory{
		Amount:      historyContext.Amount,
		PrevBalance: historyContext.PrevBalance,
		Activity:    historyContext.Activity,
		Note:        historyContext.Note,
		UserNPM:     historyContext.UserNPM,
		UserName:    historyContext.UserName,
	}

	queryAddHistory := historyContext.Db.Create(&history)
	if val, _ := validator.Query(queryAddHistory, historyContext.Context, false); !val {
		log.Printf("Error adding history | prevBal: %d, afterBal: %d, user: %s", history.PrevBalance, history.Amount, history.UserNPM)
		return true
	}
	return false
}
