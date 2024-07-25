package kassubmission

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

func GetKasSubmissions(db *gorm.DB, c *gin.Context, filters *kasSubmissionFilters) {
	var kasSubmissions []KasSubmission
	query := db.Model(&kasSubmissions)
	utils.PreloadAssociations(query, "User", "Status")
	if val, _ := validator.Query(query, c, false); !val {
		return
	}

	filterService(db, c, query, filters)
	query.Find(&kasSubmissions)
	result := make([]*KasSubmissionResponse, len(kasSubmissions))
	for i, kasSubmission := range kasSubmissions {
		result[i] = responseFormatter(&kasSubmission)
	}

	var response interface{} = result
	handler.Success(c, http.StatusOK, "Success getting kas submissions", &response, filters.Pagination)
}

func GetKasSubmissionByID(db *gorm.DB, c *gin.Context, reqUri KasSubmissionGetByID) {
	var kasSubmission KasSubmission

	query := db.Model(&kasSubmission)
	utils.PreloadAssociations(query, "User", "Status")
	query.Where("submission_id = ?", reqUri.SubmissionID)
	if val, _ := validator.Query(query, c, false); !val {
		return
	}

	query.Find(&kasSubmission)
	if kasSubmission.ID == 0 {
		handler.Error(c, http.StatusNotFound, "Kas submission not found")
		return
	}

	handler.Success(c, http.StatusOK, "Success getting kas submission", responseFormatter(&kasSubmission))
}

func CreateKasSubmission(db *gorm.DB, c *gin.Context, reqBody *KasSubmissionCreate) {
	err, userID := user.GetUserIDAndCheckNPM(db, c, reqBody.User.NPM)
	if err {
		return
	}

	var kasSubmission KasSubmission
	kasSubmission.SubmissionID = utils.GenerateRandomCapitalString(5)
	kasSubmission.UserID = userID
	kasSubmission.StatusID = consts.Pending
	kasSubmission.PayedAmount = *reqBody.PayedAmount
	kasSubmission.Note = reqBody.Note
	kasSubmission.Evidence = reqBody.Evidence

	queryCreate := db.Create(&kasSubmission)
	if val, _ := validator.Query(queryCreate, c, false); !val {
		return
	}

	queryCreated := db.Model(&kasSubmission)
	utils.PreloadAssociations(queryCreated, "User", "Status")
	queryCreated.First(&kasSubmission, "id = ?", kasSubmission.ID)
	if val, _ := validator.Query(queryCreated, c, false); !val {
		return
	}

	handler.Success(c, http.StatusCreated, "Success creating kas submission", responseFormatter(&kasSubmission))
}

func UpdateKasSubmission(db *gorm.DB, c *gin.Context, reqUri KasSubmissionGetByID, reqBody *KasSubmissionUpdate) {
	var kasSubmission KasSubmission
	query := db.Model(&kasSubmission)
	utils.PreloadAssociations(query, "User", "Status")
	query.Where("submission_id = ?", reqUri.SubmissionID)
	val, count := validator.Query(query, c, true)
	if !val {
		return
	}
	if count == 0 {
		handler.Error(c, http.StatusNotFound, "Kas submission not found")
		return
	}
	query.First(&kasSubmission)

	if err, val := enums.IsPaymentStatusValid(db, c, reqBody.Status.ID, true); err || !val {
		return
	}

	if kasSubmission.StatusID == consts.Approved || kasSubmission.StatusID == consts.Rejected {
		handler.Error(c, http.StatusBadRequest, "Kas submission already processed")
		return
	}
	if kasSubmission.StatusID == consts.Pending && reqBody.Status.ID == consts.Pending {
		handler.Error(c, http.StatusBadRequest, "Kas submission already in pending")
		return
	}
	kasSubmission.StatusID = reqBody.Status.ID

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

	if kasSubmission.StatusID == consts.Approved {

		var currentUser user.User
		queryUser := db.Model(&currentUser).Where("id = ?", kasSubmission.UserID)
		val, count := validator.Query(queryUser, c, true)
		if !val {
			return
		}
		if count == 0 {
			handler.Error(c, http.StatusBadRequest, "User not found")
			return
		}
		queryUser.First(&currentUser)
		kasSubmission.User = currentUser

		historyContext := balance.HistoryContext{
			Context:     c,
			Db:          db,
			PrevBalance: currentBalance.Balance,
			Amount:      kasSubmission.PayedAmount,
			Note:        "[kas pay] " + kasSubmission.Note,
			UserNPM:     kasSubmission.User.NPM,
			UserName:    kasSubmission.User.Name,
			Activity:    consts.Add,
		}
		if err := balance.AddHistory(historyContext); err {
			return
		}

		currentBalance.Balance += kasSubmission.PayedAmount
		queryUpdateBalance := db.Save(&currentBalance)
		if val, _ := validator.Query(queryUpdateBalance, c, false); !val {
			return
		}

		currentUser.KasPayed += kasSubmission.PayedAmount
		queryUpdatePayedKas := db.Model(&currentUser).Update("kas_payed", currentUser.KasPayed)
		if val, _ := validator.Query(queryUpdatePayedKas, c, false); !val {
			return
		}
	}

	queryUpdateKasSubmission := db.Where("submission_id = ?", reqUri.SubmissionID).
		Updates(&KasSubmission{
			StatusID: kasSubmission.StatusID,
		})
	if val, _ := validator.Query(queryUpdateKasSubmission, c, false); !val {
		return
	}

	queryUpdatedKasSubmission := db.Model(&kasSubmission)
	utils.PreloadAssociations(queryUpdatedKasSubmission, "User", "Status")
	if val, _ := validator.Query(queryUpdatedKasSubmission, c, false); !val {
		return
	}
	queryUpdatedKasSubmission.First(&kasSubmission)

	handler.Success(c, http.StatusOK, "Success updating kas submission", responseFormatter(&kasSubmission))
}
