package month

import (
	"backend/app/pkg/handler"
	"backend/app/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IsMonthExists(Db *gorm.DB, c *gin.Context, monthId int, isHandleReturn bool) (bool, bool) {
	queryMonthCheck := Db.Model(Month{}).Where("id = ?", monthId)
	val, count := validator.Query(queryMonthCheck, c, true)
	if !val {
		return true, false
	}
	if count > 0 {
		if isHandleReturn {
			handler.Error(c, http.StatusBadRequest, "Month already exists")
		}
		return false, true
	}

	return false, false
}
