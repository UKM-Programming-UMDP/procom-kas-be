package month_test

import (
	. "backend/test"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type deleteRequestParam struct {
	Year  int `form:"year"`
	Month int `form:"month"`
}
type deleteResponse struct {
	CommonResponse
	Data struct{} `json:"data"`
}

func TestDeleteMonth(t *testing.T) {
	config := InitTest()

	params := deleteRequestParam{
		Year:  2000,
		Month: 1,
	}
	CreateMonthDummy(config)
	req := Request("DELETE", WithParam("/api/month", params), nil)
	config.ServeHTTP(config.Rec, req)
	assert.Equal(t, http.StatusOK, config.Rec.Code)

	var res deleteResponse
	if err := ParseResponse(config.Rec, &res); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.Message, "Success deleting a month")
	assert.Equal(t, res.Data, struct{}{})
}
