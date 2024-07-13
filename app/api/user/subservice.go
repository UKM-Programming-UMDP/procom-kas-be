package user

import (
	"backend/app/pkg/handler"
	"backend/app/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IsNPMExists(db *gorm.DB, c *gin.Context, NPM string, isHandleReturn bool) (err bool, isExists bool) {
	queryNPMCheck := db.Model(User{}).Where("npm = ?", NPM)
	val, count := validator.Query(queryNPMCheck, c, true)
	if !val {
		return true, false
	}
	if count > 0 {
		if isHandleReturn {
			handler.Error(c, http.StatusBadRequest, "NPM already exists")
		}
		return false, true
	}

	return false, false
}

func IsEmailExists(db *gorm.DB, c *gin.Context, email string, isHandleReturn bool) (err bool, isExists bool) {
	queryEmailCheck := db.Model(User{}).Where("email = ?", email)
	val, count := validator.Query(queryEmailCheck, c, true)
	if !val {
		return true, false
	}
	if count > 0 {
		if isHandleReturn {
			handler.Error(c, http.StatusBadRequest, "Email already exists")
		}
		return false, true
	}

	return false, false
}

func GetUserIDAndCheckNPM(db *gorm.DB, c *gin.Context, NPM string) (err bool, userID int) {
	err, exist := IsNPMExists(db, c, NPM, false)
	if err {
		return true, 0
	}
	if !exist {
		handler.Error(c, http.StatusBadRequest, "User not found")
		return true, 0
	}

	var user User
	query := db.Model(&user).Where("npm = ?", NPM)
	if val, _ := validator.Query(db, c, false); !val {
		return true, 0
	}

	query.First(&user)
	return false, user.ID
}
