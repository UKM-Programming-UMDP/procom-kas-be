package user

import (
	"backend/app/api/month"
	"backend/app/common/utils"
	"backend/app/pkg/handler"
	"backend/app/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB, c *gin.Context, filters *userFilters) {
	var users []User
	query := db.Model(users).Preload("MonthStartPay")
	if val, _ := validator.Query(query, c, false); !val {
		return
	}

	filterService(c, query, filters)
	query.Find(&users)
	result := make([]*UserResponse, len(users))
	for i, user := range users {
		result[i] = responseFormatter(&user)
	}

	var response interface{} = result
	handler.Success(c, http.StatusOK, "Success getting users", &response, filters.Pagination)
}

func GetUserByNPM(db *gorm.DB, c *gin.Context, reqUri UserGetByNPM) {
	var user User
	query := db.Preload("MonthStartPay").Where("npm = ?", reqUri.NPM)
	if val, _ := validator.Query(query, c, false); !val {
		return
	}

	query.Find(&user)
	if user.ID == 0 {
		handler.Error(c, http.StatusNotFound, "User not found")
		return
	}

	response := responseFormatter(&user)
	handler.Success(c, http.StatusOK, "Success getting user", &response)
}

func CreateUser(db *gorm.DB, c *gin.Context, reqBody *UserCreate) {
	if err, exist := IsEmailExists(db, c, reqBody.Email, true); err || exist {
		return
	}
	if err, exist := IsNPMExists(db, c, reqBody.NPM, true); err || exist {
		return
	}
	err, exist := month.IsMonthExists(db, c, reqBody.MonthStartPay.ID, false)
	if err {
		return
	}
	if !exist {
		handler.Error(c, http.StatusBadRequest, "Month not found")
		return
	}

	user := User{
		NPM:           reqBody.NPM,
		Name:          reqBody.Name,
		Email:         reqBody.Email,
		KasPayed:      utils.If(reqBody.KasPayed == nil, 0, *reqBody.KasPayed),
		MonthStartPay: &reqBody.MonthStartPay,
	}

	queryCreate := db.Create(&user)
	if val, _ := validator.Query(queryCreate, c, false); !val {
		return
	}

	queryCreatedUser := db.Preload("MonthStartPay").First(&user, "npm = ?", user.NPM)
	if val, _ := validator.Query(queryCreatedUser, c, false); !val {
		return
	}

	handler.Success(c, http.StatusCreated, "Success creating user", responseFormatter(&user))
}

func UpdateUser(db *gorm.DB, c *gin.Context, reqUri UserGetByNPM, reqBody *UserUpdate) {
	var user User
	query := db.Model(&user).Where("npm = ?", reqUri.NPM)
	val, count := validator.Query(query, c, true)
	if !val {
		return
	}
	if count == 0 {
		handler.Error(c, http.StatusBadRequest, "User not found")
		return
	}

	query.First(&user)
	if user.Email != reqBody.Email {
		if err, exists := IsEmailExists(db, c, reqBody.Email, true); err || exists {
			return
		}
	}
	if user.NPM != reqUri.NPM {
		if err, exist := IsNPMExists(db, c, reqUri.NPM, true); err || exist {
			return
		}
	}
	err, exist := month.IsMonthExists(db, c, reqBody.MonthStartPay.ID, false)
	if err {
		return
	}
	if !exist {
		handler.Error(c, http.StatusBadRequest, "Month not found")
		return
	}

	user.Email = reqBody.Email
	user.NPM = reqUri.NPM
	user.Name = reqBody.Name
	user.KasPayed = *reqBody.KasPayed
	user.MonthStartPay = &reqBody.MonthStartPay

	queryUpdate := db.Save(&user)
	if val, _ := validator.Query(queryUpdate, c, false); !val {
		return
	}

	queryUpdatedUser := db.Preload("MonthStartPay").First(&user, "npm = ?", user.NPM)
	if val, _ := validator.Query(queryUpdatedUser, c, false); !val {
		return
	}

	handler.Success(c, http.StatusOK, "Success updating user", responseFormatter(&user))
}

func DeleteUser(db *gorm.DB, c *gin.Context, reqUri UserGetByNPM) {
	queryDelete := db.Where("npm = ?", reqUri.NPM).Delete(&User{})
	val, _ := validator.Query(queryDelete, c, false)
	if !val {
		return
	}
	if queryDelete.RowsAffected == 0 {
		handler.Error(c, http.StatusNotFound, "User not found")
		return
	}

	handler.Success(c, http.StatusNoContent, "", nil)
}
