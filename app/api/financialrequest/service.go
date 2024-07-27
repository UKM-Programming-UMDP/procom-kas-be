package financialrequest

import (
	"backend/app/api/balance"
	"backend/app/api/enums"
	"backend/app/api/user"
	"backend/app/common/consts"
	"backend/app/common/utils"
	"backend/app/pkg/handler"
	"backend/app/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetFinancialRequests(db *gorm.DB, c *gin.Context, filters *finreqFilters) {
	var financialRequests []FinancialRequest
	query := db.Model(&financialRequests)
	utils.PreloadAssociations(query, "User", "Status", "Type")
	if val, _ := validator.Query(query, c, false); !val {
		return
	}

	filterService(db, c, query, filters)
	query.Find(&financialRequests)
	result := make([]*FinancialRequestResponse, len(financialRequests))
	for i, financialRequest := range financialRequests {
		result[i] = responseFormatter(&financialRequest)
	}

	var response interface{} = result
	handler.Success(c, http.StatusOK, "Success getting financial requests", &response, filters.Pagination)
}

func GetFinancialRequestByID(db *gorm.DB, c *gin.Context, reqUri FinancialRequestGetByID) {
	var financialRequest FinancialRequest

	query := db.Model(&financialRequest)
	utils.PreloadAssociations(query, "User", "Status", "Type")
	query.Where("request_id = ?", reqUri.RequestID)
	if val, _ := validator.Query(query, c, false); !val {
		return
	}

	query.Find(&financialRequest)
	if financialRequest.ID == 0 {
		handler.Error(c, http.StatusNotFound, "Financial request not found")
		return
	}

	handler.Success(c, http.StatusOK, "Success getting financial request", responseFormatter(&financialRequest))
}

func CreateFinancialRequest(db *gorm.DB, c *gin.Context, reqBody *FinancialRequestCreate) {
	err, userID := user.GetUserIDAndCheckNPM(db, c, reqBody.User.NPM)
	if err {
		return
	}

	if err, val := enums.IsPaymentTypeValid(db, c, reqBody.Payment.Type.ID, true); err || !val {
		return
	}

	var finReq FinancialRequest
	finReq.RequestID = utils.GenerateRandomCapitalString(5)
	finReq.Amount = reqBody.Amount
	finReq.Note = reqBody.Note
	finReq.UserID = userID
	finReq.Payment = Payment{
		StatusID:       consts.Pending,
		Type:           reqBody.Payment.Type,
		TargetProvider: reqBody.Payment.TargetProvider,
		TargetName:     reqBody.Payment.TargetName,
		TargetNumber:   reqBody.Payment.TargetNumber,
		Evidence:       reqBody.Payment.Evidence,
	}
	finReq.TransferedEvidence = ""

	queryCreate := db.Create(&finReq)
	if val, _ := validator.Query(queryCreate, c, false); !val {
		return
	}

	queryCreatedFinreq := db.Model(&finReq)
	utils.PreloadAssociations(queryCreatedFinreq, "User", "Status", "Type")
	queryCreatedFinreq.First(&finReq, "id = ?", finReq.ID)
	if val, _ := validator.Query(queryCreatedFinreq, c, false); !val {
		return
	}

	handler.Success(c, http.StatusCreated, "Success creating financial request", responseFormatter(&finReq))
}

func UpdateFinancialRequest(db *gorm.DB, c *gin.Context, reqUri FinancialRequestGetByID, reqBody *FinancialRequestUpdate) {
	var finreq FinancialRequest
	query := db.Model(&finreq)
	utils.PreloadAssociations(query, "User", "Status", "Type")
	query.Where("request_id = ?", reqUri.RequestID)
	val, count := validator.Query(query, c, true)
	if !val {
		return
	}
	if count == 0 {
		handler.Error(c, http.StatusNotFound, "Financial request not found")
		return
	}
	query.First(&finreq)

	if err, val := enums.IsPaymentStatusValid(db, c, reqBody.Status.ID, true); err || !val {
		return
	}

	if finreq.Payment.StatusID == consts.Approved || finreq.Payment.StatusID == consts.Rejected {
		handler.Error(c, http.StatusForbidden, "Financial request already processed")
		return
	}
	if finreq.Payment.StatusID == consts.Pending && reqBody.Status.ID == consts.Pending {
		handler.Error(c, http.StatusForbidden, "Financial request already in pending")
		return
	}
	finreq.Payment.StatusID = reqBody.Status.ID

	var currentBalance balance.Balance
	queryBalance := db.Model(&currentBalance).Where("id = 1")
	val, count = validator.Query(queryBalance, c, true)
	if !val {
		return
	}
	if count == 0 {
		handler.Error(c, http.StatusNotFound, "Balance not initialized yet")
		return
	}
	queryBalance.First(&currentBalance)

	if finreq.Payment.StatusID == consts.Approved {
		var currentUser user.User
		queryUser := db.Model(currentUser).Where("id = ?", finreq.UserID)
		val, count = validator.Query(queryUser, c, true)
		if !val {
			return
		}
		if count == 0 {
			handler.Error(c, http.StatusNotFound, "User not found")
			return
		}
		queryUser.First(&currentUser)
		finreq.User = currentUser

		historyContext := balance.HistoryContext{
			Context:     c,
			Db:          db,
			PrevBalance: currentBalance.Balance,
			Amount:      finreq.Amount,
			Note:        "[finreq] " + finreq.Note,
			UserNPM:     finreq.User.NPM,
			UserName:    finreq.User.Name,
			Activity:    consts.Subtract,
		}
		if err := balance.AddHistory(historyContext); err {
			return
		}

		currentBalance.Balance -= finreq.Amount
		queryUpdateBalance := db.Save(&currentBalance)
		if val, _ := validator.Query(queryUpdateBalance, c, false); !val {
			return
		}
	}

	finreq.TransferedEvidence = reqBody.TransferedEvidence
	queryUpdateFinreq := db.Where("request_id = ?", reqUri.RequestID).
		Updates(&FinancialRequest{
			Payment: Payment{
				StatusID: reqBody.Status.ID,
			},
			TransferedEvidence: reqBody.TransferedEvidence,
		})
	if val, _ := validator.Query(queryUpdateFinreq, c, false); !val {
		return
	}

	queryUpdatedFinreq := db.Model(&finreq)
	utils.PreloadAssociations(queryUpdatedFinreq, "User", "Status", "Type")
	if val, _ := validator.Query(queryUpdatedFinreq, c, false); !val {
		return
	}
	queryUpdatedFinreq.First(&finreq)

	handler.Success(c, http.StatusOK, "Success updating financial request status", responseFormatter(&finreq))
}
