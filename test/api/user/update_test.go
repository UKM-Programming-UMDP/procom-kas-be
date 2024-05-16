package user_test

import (
	. "backend/test"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type updateRequestParam struct {
	NPM string `form:"npm"`
}

type updateRequestBody struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	KasPayed      *int   `json:"kas_payed"`
	MonthStartPay struct {
		ID int `json:"id"`
	} `json:"month_start_pay"`
}

func TestUpdateUser(t *testing.T) {
	config := InitTest()
	CreateUserDummy(config)

	params := updateRequestParam{
		NPM: "1928476912",
	}
	kasPayed := 10000
	newUser := updateRequestBody{
		Name:     "user_test_updated",
		Email:    "user_test@mail.com",
		KasPayed: &kasPayed,
		MonthStartPay: struct {
			ID int `json:"id"`
		}{ID: -1},
	}
	req := Request("PUT", WithParam("/api/users", params), newUser)
	config.ServeHTTP(config.Rec, req)
	assert.Equal(t, http.StatusOK, config.Rec.Code)
	DeleteUserDummy(config)

	type UserResponse struct {
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
	var res UserResponse
	if err := ParseResponse(config.Rec, &res); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.Message, "Success updating a user")
	assert.Equal(t, res.Data.NPM, "1928476912")
	assert.Equal(t, res.Data.Name, "user_test_updated")
	assert.Equal(t, res.Data.Email, "user_test@mail.com")
	assert.Equal(t, res.Data.KasPayed, 10000)
	assert.Equal(t, res.Data.MonthStartPay.ID, -1)
}
