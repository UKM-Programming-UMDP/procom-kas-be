package user_test

import (
	. "backend/test"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type deleteRequestParam struct {
	NPM string `form:"npm"`
}

type deleteReponse struct {
	CommonResponse
	Data struct{} `json:"data"`
}

func TestDeleteUser(t *testing.T) {
	config := InitTest()
	CreateUserDummy(config)

	params := deleteRequestParam{
		NPM: "1928476912",
	}
	req := Request("DELETE", WithParam("/api/users", params), nil)
	config.ServeHTTP(config.Rec, req)
	assert.Equal(t, http.StatusOK, config.Rec.Code)
	DeleteUserDummy(config)

	var res deleteReponse
	if err := ParseResponse(config.Rec, &res); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.Message, "Success deleting a user")
	assert.Equal(t, res.Data, struct{}{})
}
