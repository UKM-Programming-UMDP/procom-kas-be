package balance

import (
	"backend/api/balanceHistory"
	"backend/handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Db *gorm.DB
}

func (a *Controller) GetBalance(c *gin.Context) {
	balance := []BalanceSchema{}

	query := a.Db.Find(&balance)
	if val, _ := handler.QueryValidator(query, c, false); !val {
		return
	}
	if len(balance) == 0 {
		handler.Error(c, http.StatusInternalServerError, "Balance not initialized yet")
		return
	}

	handler.Success(c, http.StatusOK, "Success getting a balance", balanceResFormatter(balance[0]))
}

func (a *Controller) UpdateBalance(c *gin.Context) {
	var balance BalanceSchema
	result := a.Db.Where("id = 1").Find(&balance)
	if val, _ := handler.QueryValidator(result, c, false); !val {
		return
	}

	var reqBody BalanceUpdate
	if err := handler.BindAndValidate(c, &reqBody); err {
		return
	}

	queryCheckUser := a.Db.Where("npm = ?", reqBody.User.NPM).Find(&reqBody.User)
	if val, _ := handler.QueryValidator(queryCheckUser, c, false); !val {
		return
	}
	if reqBody.User.ID == 0 {
		handler.Error(c, http.StatusNotFound, "User not found")
		return
	}

	activity, err := balanceHistory.ActivityParser(c, reqBody.Activity)
	if err {
		return
	}

	if activity == balanceHistory.Add {
		balance.Balance += *reqBody.Amount
	} else if activity == balanceHistory.Substract {
		balance.Balance -= *reqBody.Amount
	} else {
		handler.Error(c, http.StatusBadRequest, "Invalid activity")
		return
	}

	historyContext := balanceHistory.HistoryContext{
		Context:     c,
		Db:          a.Db,
		Amount:      *reqBody.Amount,
		PrevBalance: balance.Balance,
		Note:        reqBody.Note,
		UserNPM:     reqBody.User.NPM,
		Activity:    activity,
	}
	if err := balanceHistory.AddHistory(historyContext); err {
		return
	}

	queryUpdateBalance := a.Db.Save(&balance)
	if val, _ := handler.QueryValidator(queryUpdateBalance, c, false); !val {
		return
	}

	handler.Success(c, http.StatusOK, "Success updating balance", balanceResFormatter(balance))
}

func balanceResFormatter(balance BalanceSchema) BalanceResponse {
	return BalanceResponse{
		Balance:   &balance.Balance,
		UpdatedAt: balance.UpdatedAt,
	}
}
