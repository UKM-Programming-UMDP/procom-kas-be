package kasSubmission_test

import (
	. "backend/test"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type getStatus string

const (
	getPending  getStatus = "Pending"
	getApproved getStatus = "Approved"
	getRejected getStatus = "Rejected"
)

type getResponse struct {
	CommonResponse
	Data []struct {
		SubmissionID string `json:"submission_id"`
		User         struct {
			NPM      string `json:"npm"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			KasPayed int    `json:"kas_payed"`
		} `json:"user"`
		PayedAmout  int       `json:"payed_amount"`
		Status      getStatus `json:"status"`
		Note        string    `json:"note"`
		Evidence    string    `json:"evidence"`
		SubmittedAt string    `json:"submitted_at"`
		UpdatedAt   string    `json:"updated_at"`
	} `json:"data"`
}

func TestGetAllKasSubmission(t *testing.T) {
	config := InitTest()

	CreateKasSubmissionDummy(config)
	req := Request("GET", "/api/kas", nil)
	config.ServeHTTP(config.Rec, req)
	assert.Equal(t, http.StatusOK, config.Rec.Code)
	DeleteKasSubmissionDummy(config)

	var res getResponse
	if err := ParseResponse(config.Rec, &res); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.Message, "Success getting kas submissions")
	assert.Len(t, res.Data[0].SubmissionID, 5)
	assert.Len(t, res.Data[0].User.NPM, 10)
	assert.NotEmpty(t, res.Data[0].User.Name)
	assert.NotEmpty(t, res.Data[0].User.Email)
	assert.GreaterOrEqual(t, res.Data[0].User.KasPayed, 0)
	assert.GreaterOrEqual(t, res.Data[0].PayedAmout, 0)
	assert.NotEmpty(t, res.Data[0].Status)
	assert.Contains(t, []getStatus{getPending, getApproved, getRejected}, res.Data[0].Status)
	assert.NotEmpty(t, res.Data[0].Note)
	assert.NotEmpty(t, res.Data[0].Evidence)
	assert.NotEmpty(t, res.Data[0].SubmittedAt)
	assert.NotEmpty(t, res.Data[0].UpdatedAt)
}
