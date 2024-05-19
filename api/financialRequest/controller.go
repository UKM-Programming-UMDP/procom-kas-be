package financialRequest

import (
	"backend/api/balance"
	"backend/api/balanceHistory"
	"backend/api/user"
	"backend/handler"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Db *gorm.DB
}

func (a *Controller) GetFinancialRequest(c *gin.Context) {
	var financialRequests []FinancialRequestSchema
	query := a.Db.Preload("User").Model(financialRequests)
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

	query.Find(&financialRequests)
	result := make([]FinancialRequestResponse, len(financialRequests))
	for i, finReq := range financialRequests {
		result[i] = finReqResponseFormatter(finReq)
	}

	handler.Success(c, http.StatusOK, "Success getting financial requests", result, pagination)
}

func (a *Controller) GetFinancialRequestById(c *gin.Context) {
	var reqParam FinancialRequestParam
	if err := handler.BindParamAndValidate(c, &reqParam); err {
		return
	}

	var finReq FinancialRequestSchema
	query := a.Db.Preload("User").Where("request_id = ?", reqParam.RequestID).Find(&finReq)
	if val, _ := handler.QueryValidator(query, c, false); !val {
		return
	}
	if finReq.ID == 0 {
		handler.Error(c, http.StatusNotFound, "Financial request not found")
		return
	}

	handler.Success(c, http.StatusOK, "Success getting a financial request", finReqResponseFormatter(finReq))
}

func (a *Controller) CreateFinancialRequest(c *gin.Context) {
	var reqBody FinancialRequestCreate
	if err := handler.BindAndValidate(c, &reqBody); err {
		return
	}

	var finReq FinancialRequestSchema
	queryCheckUser := a.Db.Where("npm = ?", reqBody.User.NPM).Find(&finReq.User)
	if val, _ := handler.QueryValidator(queryCheckUser, c, false); !val {
		return
	}
	if finReq.User.ID == 0 {
		handler.Error(c, http.StatusNotFound, "User not found")
		return
	}

	finReq.RequestID = utils.RandomString(5)
	finReq.Amount = reqBody.Amount
	finReq.Note = reqBody.Note
	finReq.Status = Pending
	finReq.Payment = Payment{
		Type:           reqBody.Payment.Type,
		TargetProvider: reqBody.Payment.TargetProvider,
		TargetName:     reqBody.Payment.TargetName,
		TargetNumber:   reqBody.Payment.TargetNumber,
		Evidence:       reqBody.Payment.Evidence,
	}
	finReq.TransferedEvidence = ""

	queryCreate := a.Db.Create(&finReq)
	if val, _ := handler.QueryValidator(queryCreate, c, false); !val {
		return
	}

	handler.Success(c, http.StatusCreated, "Success creating a financial request", finReqResponseFormatter(finReq))
}

func (a *Controller) UpdateFinancialRequestStatus(c *gin.Context) {
	var reqParam FinancialRequestParam
	if err := handler.BindParamAndValidate(c, &reqParam); err {
		return
	}
	var reqBody FinancialRequestUpdate
	if err := handler.BindAndValidate(c, &reqBody); err {
		return
	}

	var finReq FinancialRequestSchema
	query := a.Db.Preload("User").Where("request_id = ?", reqParam.RequestID).Find(&finReq)
	if val, _ := handler.QueryValidator(query, c, false); !val {
		return
	}
	if finReq.ID == 0 {
		handler.Error(c, http.StatusNotFound, "Financial request not found")
		return
	}
	if finReq.Status == Approved || finReq.Status == Rejected {
		handler.Error(c, http.StatusConflict, "Financial request already processed")
		return
	}
	if finReq.Status == Pending && *reqBody.Status == 3 {
		handler.Error(c, http.StatusBadRequest, "Cannot change status to pending")
		return
	}
	if err := StatusParser(c, reqBody.Status, &finReq.Status); err {
		return
	}

	var balance balance.BalanceSchema
	queryBalance := a.Db.Where("id = 1").Find(&balance)
	if val, _ := handler.QueryValidator(queryBalance, c, false); !val {
		return
	}

	if finReq.Status == Approved {
		var currentUser user.UserSchema
		queryGetUser := a.Db.Where("npm = ?", finReq.User.NPM).Find(&currentUser)
		if val, _ := handler.QueryValidator(queryGetUser, c, false); !val {
			return
		}
		if currentUser.ID == 0 {
			handler.Error(c, http.StatusNotFound, "User not found")
			return
		}
		finReq.User = &currentUser

		historyContext := balanceHistory.HistoryContext{
			Context:     c,
			Db:          a.Db,
			PrevBalance: balance.Balance,
			Amount:      finReq.Amount,
			Note:        "[finreq] " + finReq.Note,
			UserNPM:     finReq.UserNPM,
			UserName:    finReq.User.Name,
			Activity:    balanceHistory.Substract,
		}
		if err := balanceHistory.AddHistory(historyContext); err {
			return
		}

		balance.Balance -= finReq.Amount
		queryUpdateBalance := a.Db.Save(&balance)
		if val, _ := handler.QueryValidator(queryUpdateBalance, c, false); !val {
			return
		}
	}

	finReq.TransferedEvidence = reqBody.TransferedEvidence
	queryUpdate := a.Db.Save(&finReq)
	if val, _ := handler.QueryValidator(queryUpdate, c, false); !val {
		return
	}

	handler.Success(c, http.StatusOK, "Success updating financial request status", finReqResponseFormatter(finReq))
}

func StatusParser(c *gin.Context, activity *int, statusTarget *status) bool {
	statusMap := map[int]status{
		1: Approved,
		2: Rejected,
		3: Pending,
	}
	if _, ok := statusMap[*activity]; !ok {
		handler.Error(c, http.StatusBadRequest, "Invalid status")
		return true
	}
	*statusTarget = statusMap[*activity]
	return false
}

func finReqResponseFormatter(finReq FinancialRequestSchema) FinancialRequestResponse {
	return FinancialRequestResponse{
		RequestID: finReq.RequestID,
		Amount:    finReq.Amount,
		Note:      finReq.Note,
		User: user.UserResponse{
			NPM:   finReq.User.NPM,
			Name:  finReq.User.Name,
			Email: finReq.User.Email,
		},
		Status:    finReq.Status,
		Payment:   finReq.Payment,
		CreatedAt: finReq.CreatedAt,
		UpdatedAt: finReq.UpdatedAt,
	}
}
