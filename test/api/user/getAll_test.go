package user_test

import (
	. "backend/test"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type getAllResponse struct {
	CommonResponse
	Data []struct {
		NPM           string `json:"npm"`
		Name          string `json:"name"`
		Email         string `json:"email"`
		KasPayed      int    `json:"kas_payed"`
		MonthStartPay struct {
			ID int `json:"id"`
		} `json:"month_start_pay"`
	} `json:"data"`
}

func TestGetAllUser(t *testing.T) {
	config := InitTest()
	DeleteUserDummy(config)

	CreateUserDummy(config)
	req := Request("GET", "/api/users", nil)
	config.ServeHTTP(config.Rec, req)
	assert.Equal(t, http.StatusOK, config.Rec.Code)
	DeleteUserDummy(config)

	var res getAllResponse
	if err := ParseResponse(config.Rec, &res); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Success getting users", res.Message)
	assert.Len(t, res.Data[0].NPM, 10)
	assert.NotEmpty(t, res.Data[0].Name)
	assert.NotEmpty(t, res.Data[0].Email)
	assert.GreaterOrEqual(t, res.Data[0].KasPayed, 0)
	assert.NotNil(t, res.Data[0].MonthStartPay.ID)
}
