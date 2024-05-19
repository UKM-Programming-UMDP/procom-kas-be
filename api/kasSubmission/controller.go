package kasSubmission

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

func (a *Controller) GetKasSubmissions(c *gin.Context) {
	var kasSubmissions []KasSubmissionSchema
	query := a.Db.Model(kasSubmissions).Preload("User")
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

	query.Find(&kasSubmissions)
	result := make([]KasSubmissionResponse, len(kasSubmissions))
	for i, kasSubmission := range kasSubmissions {
		result[i] = kasResFormatter(kasSubmission)
	}

	handler.Success(c, http.StatusOK, "Success getting kas submissions", result, pagination)
}

func (a *Controller) GetKasSubmissionById(c *gin.Context) {
	var reqParam KasSubmissionRequestParam
	if err := handler.BindParamAndValidate(c, &reqParam); err {
		return
	}

	var kasSubmission KasSubmissionSchema
	query := a.Db.Preload("User").Where("submission_id = ?", reqParam.SubmissionID).Find(&kasSubmission)
	if val, _ := handler.QueryValidator(query, c, false); !val {
		return
	}
	if kasSubmission.ID == 0 {
		handler.Error(c, http.StatusNotFound, "Kas submission not found")
		return
	}

	handler.Success(c, http.StatusOK, "Success getting a kas submission", kasResFormatter(kasSubmission))
}

func (a *Controller) CreateKasSubmission(c *gin.Context) {
	var reqBody KasSubmissionCreate
	if err := handler.BindAndValidate(c, &reqBody); err {
		return
	}

	var kasSubmission KasSubmissionSchema
	queryCheckUser := a.Db.Where("npm = ?", reqBody.User.NPM).Find(&kasSubmission.User)
	if val, _ := handler.QueryValidator(queryCheckUser, c, false); !val {
		return
	}
	if kasSubmission.User.ID == 0 {
		handler.Error(c, http.StatusNotFound, "User not found")
		return
	}

	kasSubmission.Status = Pending
	kasSubmission.SubmissionID = utils.RandomString(5)
	kasSubmission.PayedAmount = utils.If(reqBody.PayedAmount == nil, 0, *reqBody.PayedAmount)
	kasSubmission.Note = utils.If(reqBody.Note == nil, "", *reqBody.Note)
	kasSubmission.Evidence = reqBody.Evidence

	queryCreate := a.Db.Create(&kasSubmission)
	if val, _ := handler.QueryValidator(queryCreate, c, false); !val {
		return
	}

	handler.Success(c, http.StatusCreated, "Success creating a kas submission", kasResFormatter(kasSubmission))
}

func (a *Controller) UpdateKasSubmissionStatus(c *gin.Context) {
	var reqParam KasSubmissionRequestParam
	if err := handler.BindParamAndValidate(c, &reqParam); err {
		return
	}
	var reqBody KasSubmissionUpdateStatus
	if err := handler.BindAndValidate(c, &reqBody); err {
		return
	}

	var kasSubmission KasSubmissionSchema
	query := a.Db.Preload("User").Where("submission_id = ?", reqParam.SubmissionID).Find(&kasSubmission)
	if val, _ := handler.QueryValidator(query, c, false); !val {
		return
	}
	if kasSubmission.ID == 0 {
		handler.Error(c, http.StatusNotFound, "Kas submission not found")
		return
	}
	if kasSubmission.Status == Approved || kasSubmission.Status == Rejected {
		handler.Error(c, http.StatusBadRequest, "Kas submission already processed")
		return
	}
	if kasSubmission.Status == Pending && *reqBody.Status == 3 {
		handler.Error(c, http.StatusBadRequest, "Cannot change status to pending")
		return
	}
	if err := StatusParser(c, reqBody.Status, &kasSubmission.Status); err {
		return
	}

	if kasSubmission.Status == Approved {
		var currentUser user.UserSchema
		queryGetUser := a.Db.Where("npm = ?", kasSubmission.User.NPM).Find(&currentUser)
		if val, _ := handler.QueryValidator(queryGetUser, c, false); !val {
			return
		}
		if currentUser.ID == 0 {
			handler.Error(c, http.StatusNotFound, "User not found")
			return
		}
		currentUser.KasPayed += kasSubmission.PayedAmount

		var balance balance.BalanceSchema
		queryBalance := a.Db.Where("id = 1").Find(&balance)
		if val, _ := handler.QueryValidator(queryBalance, c, false); !val {
			return
		}

		historyContext := balanceHistory.HistoryContext{
			Context:     c,
			Db:          a.Db,
			PrevBalance: balance.Balance,
			Amount:      kasSubmission.PayedAmount,
			Note:        "[kas pay]",
			UserNPM:     kasSubmission.User.NPM,
			UserName:    kasSubmission.User.Name,
			Activity:    balanceHistory.Add,
		}
		if err := balanceHistory.AddHistory(historyContext); err {
			return
		}

		balance.Balance += kasSubmission.PayedAmount
		queryUpdateBalance := a.Db.Save(&balance)
		if val, _ := handler.QueryValidator(queryUpdateBalance, c, false); !val {
			return
		}

		updateUserKasPayedQuery := a.Db.Model(&currentUser).Update("kas_payed", currentUser.KasPayed)
		if val, _ := handler.QueryValidator(updateUserKasPayedQuery, c, false); !val {
			return
		}
	}

	updateQuery := a.Db.Model(&kasSubmission).Clauses(utils.Returning()).Update("status", kasSubmission.Status)
	if val, _ := handler.QueryValidator(updateQuery, c, false); !val {
		return
	}

	handler.Success(c, http.StatusOK, "Success updating kas submission status", kasResFormatter(kasSubmission))
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

func kasResFormatter(kasSubmission KasSubmissionSchema) KasSubmissionResponse {
	return KasSubmissionResponse{
		SubmissionID: kasSubmission.SubmissionID,
		User: user.UserResponse{
			NPM:      kasSubmission.User.NPM,
			Name:     kasSubmission.User.Name,
			Email:    kasSubmission.User.Email,
			KasPayed: &kasSubmission.User.KasPayed,
		},
		PayedAmount: &kasSubmission.PayedAmount,
		Status:      kasSubmission.Status,
		Note:        &kasSubmission.Note,
		Evidence:    kasSubmission.Evidence,
		SubmittedAt: kasSubmission.CreatedAt,
		UpdatedAt:   kasSubmission.UpdatedAt,
	}
}
