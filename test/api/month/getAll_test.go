package month_test

import (
	. "backend/test"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type getAllReponse struct {
	CommonResponse
	Data []struct {
		ID    int  `json:"id"`
		Year  int  `json:"year"`
		Month int  `json:"month"`
		Show  bool `json:"show"`
	} `json:"data"`
}

func TestGetAllMonth(t *testing.T) {
	config := InitTest()

	CreateMonthDummy(config)
	req := Request("GET", "/api/month", nil)
	config.ServeHTTP(config.Rec, req)
	assert.Equal(t, http.StatusOK, config.Rec.Code)
	DeleteMonthDummy(config)

	var res getAllReponse
	if err := ParseResponse(config.Rec, &res); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.Message, "Success getting a month")
	assert.GreaterOrEqual(t, res.Data[0].Year, 2000)
	assert.LessOrEqual(t, res.Data[0].Year, 9999)
	assert.GreaterOrEqual(t, res.Data[0].Month, 1)
	assert.LessOrEqual(t, res.Data[0].Month, 12)
	assert.Equal(t, false, res.Data[0].Show)
}
