package balance

import (
	"backend/app/common/consts"
	"backend/app/pkg/handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HistoryContext struct {
	Context     *gin.Context
	Db          *gorm.DB
	Amount      int
	PrevBalance int
	Note        string
	UserNPM     string
	UserName    string
	Activity    consts.Activity
}

func activityParser(c *gin.Context, activity *int) (consts.Activity, bool) {
	activityMap := map[int]consts.Activity{
		1: consts.Add,
		2: consts.Subtract,
	}
	if _, ok := activityMap[*activity]; !ok {
		handler.Error(c, http.StatusBadRequest, "Invalid activity")
		return "", true
	}
	return activityMap[*activity], false
}

func balanceResFormatter(balance *Balance) *BalanceResponse {
	return &BalanceResponse{
		Balance:   &balance.Balance,
		UpdatedAt: balance.UpdatedAt,
	}
}

func historyResFormatter(history *BalanceHistory) *BalanceHistoryResponse {
	return &BalanceHistoryResponse{
		Amount:      history.Amount,
		PrevBalance: history.PrevBalance,
		Activity:    history.Activity,
		Note:        history.Note,
		User: struct {
			NPM  string `json:"npm,omitempty"`
			Name string `json:"name,omitempty"`
		}{
			NPM:  history.UserNPM,
			Name: history.UserName,
		},
		CreatedAt: history.CreatedAt,
	}
}
