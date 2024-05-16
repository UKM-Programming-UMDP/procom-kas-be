package month_test

import (
	. "backend/test"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type updateRequestParam struct {
	Year  int `form:"year"`
	Month int `form:"month"`
}

type updateRequestBody struct {
	Show *bool `json:"show"`
}

type updateResponse struct {
	CommonResponse
	Data struct {
		ID    int  `json:"id"`
		Year  int  `json:"year"`
		Month int  `json:"month"`
		Show  bool `json:"show"`
	} `json:"data"`
}

func TestUpdateMonth(t *testing.T) {
	config := InitTest()

	params := updateRequestParam{
		Year:  2000,
		Month: 1,
	}
	var show = true
	newMonth := updateRequestBody{
		Show: &show,
	}
	CreateMonthDummy(config)
	req := Request("PUT", WithParam("/api/month", params), newMonth)
	config.ServeHTTP(config.Rec, req)
	assert.Equal(t, http.StatusOK, config.Rec.Code)
	DeleteMonthDummy(config)

	var res updateResponse
	if err := ParseResponse(config.Rec, &res); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.Message, "Success updating show month")
	assert.Equal(t, res.Data.Year, 2000)
	assert.Equal(t, res.Data.Month, 1)
	assert.Equal(t, true, res.Data.Show)
}
