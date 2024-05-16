package user

import (
	"backend/api/month"
	"backend/handler"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Db *gorm.DB
}

func (a *Controller) GetUsers(c *gin.Context) {
	var users []UserSchema
	query := a.Db.Model(users).Preload("MonthStartPay")
	val, count := handler.QueryValidator(query, c, true)
	if !val {
		return
	}
	pagination, err := handler.PaginationBuilder(c, query, &count)
	if err {
		return
	}
	if err := handler.FilterBuilder(c, query, "name"); err {
		return
	}

	query.Find(&users)
	result := make([]UserResponse, len(users))
	for i, user := range users {
		result[i] = userResFormatter(user)
	}

	handler.Success(c, http.StatusOK, "Success getting users", result, pagination)
}

func (a *Controller) GetUserById(c *gin.Context) {
	var reqUser UserRequestParam
	if err := handler.BindParamAndValidate(c, &reqUser); err {
		return
	}

	var user UserSchema
	query := a.Db.Preload("MonthStartPay").Where("NPM = ?", reqUser.NPM).Find(&user)
	if val, _ := handler.QueryValidator(query, c, false); !val {
		return
	}

	if user.ID == 0 {
		handler.Error(c, http.StatusNotFound, "User not found")
		return
	}

	handler.Success(c, http.StatusOK, "Success getting a user", userResFormatter(user))
}

func (a *Controller) CreateUser(c *gin.Context) {
	var reqBody UserCreate
	if err := handler.BindAndValidate(c, &reqBody); err {
		return
	}

	queryNpmCheck := a.Db.Model(UserSchema{}).Where("npm = ?", reqBody.NPM)
	val, count := handler.QueryValidator(queryNpmCheck, c, true)
	if !val {
		return
	}
	if count != 0 {
		handler.Error(c, http.StatusConflict, "User already exists")
		return
	}

	queryEmailCheck := a.Db.Model(UserSchema{}).Where("email = ?", reqBody.Email)
	val, count = handler.QueryValidator(queryEmailCheck, c, true)
	if !val {
		return
	}
	if count != 0 {
		handler.Error(c, http.StatusConflict, "Email already used")
		return
	}

	queryMonthIdCheck := a.Db.Model(month.MonthSchema{}).Where("id = ?", reqBody.MonthStartPay.ID)
	val, count = handler.QueryValidator(queryMonthIdCheck, c, true)
	if !val {
		return
	}
	if count == 0 {
		handler.Error(c, http.StatusNotFound, "Month not found or registered yet")
		return
	}

	user := UserSchema{
		NPM:           reqBody.NPM,
		Name:          reqBody.Name,
		Email:         reqBody.Email,
		KasPayed:      utils.If(reqBody.KasPayed == nil, 0, *reqBody.KasPayed),
		MonthStartPay: &reqBody.MonthStartPay,
	}

	queryCreate := a.Db.Create(&user)
	if val, _ := handler.QueryValidator(queryCreate, c, false); !val {
		return
	}

	handler.Success(c, http.StatusCreated, "Success creating a user", userResFormatter(user))
}

func (a *Controller) UpdateUser(c *gin.Context) {
	var reqParam UserRequestParam
	if err := handler.BindParamAndValidate(c, &reqParam); err {
		return
	}

	var prevUser UserSchema
	queryUser := a.Db.Preload("MonthStartPay").Where("NPM = ?", reqParam.NPM).Find(&prevUser)
	if val, _ := handler.QueryValidator(queryUser, c, false); !val {
		return
	}
	if prevUser.ID == 0 {
		handler.Error(c, http.StatusNotFound, "User not found")
		return
	}

	var reqBody UserUpdate
	if err := handler.BindAndValidate(c, &reqBody); err {
		return
	}

	if prevUser.Email != reqBody.Email {
		var tempUser UserSchema
		queryDuplicateCheck := a.Db.Where("email = ? and npm != ?", reqBody.Email, prevUser.NPM).Clauses(utils.Returning("ID")).Find(&tempUser)
		if val, _ := handler.QueryValidator(queryDuplicateCheck, c, false); !val {
			return
		}
		if tempUser.ID != 0 {
			handler.Error(c, http.StatusConflict, "Email already used")
			return
		}
	}

	queryMonthDupCheck := a.Db.Model(month.MonthSchema{}).Where("id = ?", reqBody.MonthStartPay.ID)
	val, count := handler.QueryValidator(queryMonthDupCheck, c, true)
	if !val {
		return
	}
	if count == 0 {
		handler.Error(c, http.StatusNotFound, "Month not found")
		return
	}

	prevUser.Name = reqBody.Name
	prevUser.Email = reqBody.Email
	prevUser.KasPayed = *reqBody.KasPayed
	prevUser.MonthStartPay = &reqBody.MonthStartPay

	queryUpdate := a.Db.Clauses(utils.Returning()).Where("npm = ?", reqParam.NPM).Updates(&prevUser)
	if val, _ := handler.QueryValidator(queryUpdate, c, false); !val {
		return
	}

	handler.Success(c, http.StatusOK, "Success updating a user", userResFormatter(prevUser))
}

func (a *Controller) DeleteUser(c *gin.Context) {
	var reqParam UserRequestParam
	if err := handler.BindParamAndValidate(c, &reqParam); err {
		return
	}

	var user UserDelete
	queryDelete := a.Db.Model(UserSchema{}).Clauses(utils.Returning("id")).Where("NPM = ?", reqParam.NPM).Delete(&user)
	if val, _ := handler.QueryValidator(queryDelete, c, false); !val {
		return
	}

	if user.ID == 0 {
		handler.Error(c, http.StatusNotFound, "User not found")
		return
	}

	handler.Success(c, http.StatusOK, "Success deleting a user", nil)
}

func CheckIfUserExist(db *gorm.DB, npm string) (err bool, isExist bool) {
	var user UserSchema
	query := db.Where("npm = ?", npm).Find(&user)
	val, count := handler.QueryValidator(query, nil, true)
	if !val {
		return true, false
	}
	if count == 0 {
		return false, false
	}
	return false, true
}

func userResFormatter(user UserSchema) UserResponse {
	return UserResponse{
		NPM:      user.NPM,
		Name:     user.Name,
		Email:    user.Email,
		KasPayed: &user.KasPayed,
		MonthStartPay: &month.MonthResponse{
			ID: user.MonthStartPay.ID,
		},
	}
}
