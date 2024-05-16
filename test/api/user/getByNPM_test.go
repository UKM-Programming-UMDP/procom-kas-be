package user_test

import (
	. "backend/test"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type getByNPMRequestParam struct {
	NPM string `form:"npm"`
}

type getByNPMResponse struct {
	CommonResponse
	Data struct {
		NPM           string `json:"npm"`
		Name          string `json:"name"`
		Email         string `json:"email"`
		KasPayed      int    `json:"kas_payed"`
		MonthStartPay struct {
			ID int `json:"id"`
		} `json:"month_start_pay"`
	} `json:"data"`
}

func TestGetUserByNPM(t *testing.T) {
	config := InitTest()

	CreateUserDummy(config)
	params := getByNPMRequestParam{
		NPM: "1928476912",
	}
	req := Request("GET", WithParam("/api/users/details", params), nil)
	config.ServeHTTP(config.Rec, req)
	assert.Equal(t, http.StatusOK, config.Rec.Code)
	DeleteUserDummy(config)

	var res getByNPMResponse
	if err := ParseResponse(config.Rec, &res); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.Message, "Success getting a user")
	assert.Equal(t, res.Data.NPM, "1928476912")
	assert.Equal(t, res.Data.Name, "user_test")
	assert.Equal(t, res.Data.Email, "user_test@mail.com")
	assert.Equal(t, res.Data.KasPayed, 0)
	assert.Equal(t, res.Data.MonthStartPay.ID, -1)
}
