package kassubmission

import (
	"backend/app/api/user"
	"backend/app/pkg/validator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Controller(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/v1")

	v1.GET("/kas-submissions", func(c *gin.Context) {
		var filters kasSubmissionFilters
		if err := validator.BindParams(c, &filters); err {
			return
		}

		GetKasSubmissions(db, c, &filters)
	})

	v1.GET("/kas-submissions/:id", func(c *gin.Context) {
		var reqUri KasSubmissionGetByID
		if err := validator.BindUri(c, &reqUri); err {
			return
		}

		GetKasSubmissionByID(db, c, reqUri)
	})

	v1.POST("/kas-submissions", func(c *gin.Context) {
		var reqBody KasSubmissionCreate
		if err := validator.BindBody(c, &reqBody); err {
			return
		}

		CreateKasSubmission(db, c, &reqBody)
	})

	v1.PUT("/kas-submissions/:id", func(c *gin.Context) {
		var reqUri KasSubmissionGetByID
		if err := validator.BindUri(c, &reqUri); err {
			return
		}

		var reqBody KasSubmissionUpdate
		if err := validator.BindBody(c, &reqBody); err {
			return
		}

		UpdateKasSubmission(db, c, reqUri, &reqBody)
	})
}

func responseFormatter(kasSubmission *KasSubmission) *KasSubmissionResponse {
	return &KasSubmissionResponse{
		SubmissionID: kasSubmission.SubmissionID,
		User: user.UserResponse{
			NPM:      kasSubmission.User.NPM,
			Name:     kasSubmission.User.Name,
			Email:    kasSubmission.User.Email,
			KasPayed: &kasSubmission.User.KasPayed,
		},
		PayedAmount: &kasSubmission.PayedAmount,
		Status:      kasSubmission.Status,
		Note:        kasSubmission.Note,
		Evidence:    kasSubmission.Evidence,
		SubmittedAt: kasSubmission.CreatedAt,
		UpdatedAt:   kasSubmission.UpdatedAt,
	}
}
