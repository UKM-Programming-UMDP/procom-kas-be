package user_test

import (
	. "backend/test"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type createRequestBody struct {
	NPM           string `json:"npm"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	KasPayed      *int   `json:"kas_payed"`
	MonthStartPay struct {
		ID int `json:"id"`
	} `json:"month_start_pay"`
}

type createReponse struct {
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

func TestCreateUser(t *testing.T) {
	config := InitTest()
	CreateMonthDummy(config)

	newUser := createRequestBody{
		NPM:      "1928476912",
		Name:     "user_test",
		Email:    "user_test@mail.com",
		KasPayed: new(int),
		MonthStartPay: struct {
			ID int `json:"id"`
		}{ID: -1},
	}
	req := Request("POST", "/api/users", newUser)
	config.ServeHTTP(config.Rec, req)
	assert.Equal(t, http.StatusCreated, config.Rec.Code)
	DeleteUserDummy(config)

	var res createReponse
	if err := ParseResponse(config.Rec, &res); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.Message, "Success creating a user")
	assert.Equal(t, res.Data.NPM, "1928476912")
	assert.Equal(t, res.Data.Name, "user_test")
	assert.Equal(t, res.Data.Email, "user_test@mail.com")
	assert.Equal(t, res.Data.KasPayed, 0)
	assert.Equal(t, res.Data.MonthStartPay.ID, -1)
}
