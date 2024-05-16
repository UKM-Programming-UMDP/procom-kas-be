package month_test

import (
	. "backend/test"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type createRequestParam struct {
	Year  int `form:"year"`
	Month int `form:"month"`
}
type createResponse struct {
	CommonResponse
	Data struct {
		ID    int  `json:"id"`
		Year  int  `json:"year"`
		Month int  `json:"month"`
		Show  bool `json:"show"`
	} `json:"data"`
}

func TestCreateMonth(t *testing.T) {
	config := InitTest()

	newMonth := createRequestParam{
		Year:  2000,
		Month: 1,
	}
	req := Request("POST", "/api/month", newMonth)
	config.ServeHTTP(config.Rec, req)
	assert.Equal(t, http.StatusCreated, config.Rec.Code)
	DeleteMonthDummy(config)

	var res createResponse
	if err := ParseResponse(config.Rec, &res); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.Message, "Success creating a month")
	assert.Equal(t, res.Data.Year, 2000)
	assert.Equal(t, res.Data.Month, 1)
	assert.Equal(t, false, res.Data.Show)
}
