package user

import (
	"backend/app/api/month"
	"backend/app/pkg/validator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Controller(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/v1")

	v1.GET("/users", func(c *gin.Context) {
		var filters userFilters
		if err := validator.BindParams(c, &filters); err {
			return
		}

		GetUsers(db, c, &filters)
	})

	v1.GET("/users/:npm", func(c *gin.Context) {
		var reqUri UserGetByNPM
		if err := validator.BindUri(c, &reqUri); err {
			return
		}

		GetUserByNPM(db, c, reqUri)
	})

	v1.POST("/users", func(c *gin.Context) {
		var reqBody UserCreate
		if err := validator.BindBody(c, &reqBody); err {
			return
		}

		CreateUser(db, c, &reqBody)
	})

	v1.PUT("/users/:npm", func(c *gin.Context) {
		var reqUri UserGetByNPM
		if err := validator.BindUri(c, &reqUri); err {
			return
		}

		var reqBody UserUpdate
		if err := validator.BindBody(c, &reqBody); err {
			return
		}

		UpdateUser(db, c, reqUri, &reqBody)
	})

	v1.DELETE("/users/:npm", func(c *gin.Context) {
		var reqUri UserGetByNPM
		if err := validator.BindUri(c, &reqUri); err {
			return
		}

		DeleteUser(db, c, reqUri)
	})

}

func responseFormatter(user *User) *UserResponse {
	return &UserResponse{
		NPM:      user.NPM,
		Name:     user.Name,
		Email:    user.Email,
		KasPayed: &user.KasPayed,
		MonthStartPay: &month.MonthResponse{
			ID: user.MonthStartPay.ID,
		},
	}
}
