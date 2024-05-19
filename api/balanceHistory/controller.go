package balanceHistory

import (
	"backend/handler"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Db *gorm.DB
}

func (a *Controller) GetHistory(c *gin.Context) {
	var balances []BalanceHistorySchema
	query := a.Db.Model(balances)
	val, count := handler.QueryValidator(query, c, true)
	if !val {
		return
	}
	pagination, err := handler.PaginationBuilder(c, query, &count)
	if err {
		return
	}
	if err := handler.FilterBuilder(c, query); err {
		return
	}

	query.Find(&balances)
	result := make([]BalanceHistoryResponse, len(balances))
	for i, balance := range balances {
		result[i] = balanceResFormatter(balance)
	}

	handler.Success(c, http.StatusOK, "Success getting balance history", result, pagination)
}

type HistoryContext struct {
	Context     *gin.Context
	Db          *gorm.DB
	Amount      int
	PrevBalance int
	Note        string
	UserNPM     string
	UserName    string
	Activity    Activity
}

func AddHistory(historyContext HistoryContext) bool {
	history := BalanceHistorySchema{
		Amount:      historyContext.Amount,
		PrevBalance: historyContext.PrevBalance,
		Activity:    historyContext.Activity,
		Note:        historyContext.Note,
		UserNPM:     historyContext.UserNPM,
		UserName:    historyContext.UserName,
	}

	queryAddHistory := historyContext.Db.Create(&history)
	if val, _ := handler.QueryValidator(queryAddHistory, historyContext.Context, false); !val {
		log.Printf("Error adding history | prevBal: %d, afterBal: %d, user: %s", history.PrevBalance, history.Amount, history.UserNPM)
		return true
	}
	return false
}

func ActivityParser(c *gin.Context, activity *int) (Activity, bool) {
	activityMap := map[int]Activity{
		1: Add,
		2: Substract,
	}
	if _, ok := activityMap[*activity]; !ok {
		handler.Error(c, http.StatusBadRequest, "Invalid activity")
		return "", true
	}
	return activityMap[*activity], false
}

func balanceResFormatter(balance BalanceHistorySchema) BalanceHistoryResponse {
	return BalanceHistoryResponse{
		Amount:      balance.Amount,
		PrevBalance: balance.PrevBalance,
		Activity:    balance.Activity,
		Note:        balance.Note,
		User: struct {
			NPM  string `json:"npm,omitempty"`
			Name string `json:"name,omitempty"`
		}{
			NPM:  balance.UserNPM,
			Name: balance.UserName,
		},
		CreatedAt: balance.CreatedAt,
	}
}
