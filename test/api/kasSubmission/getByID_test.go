package kasSubmission_test

import (
	. "backend/test"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type getByIDRequestParam struct {
	SubmissionID string `form:"submission_id"`
}

type getByIDStatus string

const (
	getByIDPending  getByIDStatus = "Pending"
	getByIDApproved getByIDStatus = "Approved"
	getByIDRejected getByIDStatus = "Rejected"
)

type KasSubmissionResponse struct {
	CommonResponse
	Data struct {
		SubmissionID string `json:"submission_id"`
		User         struct {
			NPM      string `json:"npm"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			KasPayed int    `json:"kas_payed"`
		} `json:"user"`
		PayedAmout  int           `json:"payed_amount"`
		Status      getByIDStatus `json:"status"`
		Note        string        `json:"note"`
		Evidence    string        `json:"evidence"`
		SubmittedAt string        `json:"submitted_at"`
		UpdatedAt   string        `json:"updated_at"`
	} `json:"data"`
}

func TestGetKasSubmissionByID(t *testing.T) {
	config := InitTest()

	CreateKasSubmissionDummy(config)
	params := getByIDRequestParam{
		SubmissionID: "test1",
	}
	req := Request("GET", WithParam("/api/kas/details", params), nil)
	config.ServeHTTP(config.Rec, req)
	assert.Equal(t, http.StatusOK, config.Rec.Code)
	DeleteKasSubmissionDummy(config)

	var res KasSubmissionResponse
	if err := ParseResponse(config.Rec, &res); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.Message, "Success getting a kas submission")
	assert.Equal(t, res.Data.SubmissionID, "test1")
	assert.Equal(t, res.Data.User.NPM, "1928476912")
	assert.Equal(t, res.Data.User.Name, "user_test")
	assert.Equal(t, res.Data.User.Email, "user_test@mail.com")
	assert.Equal(t, res.Data.User.KasPayed, 0)
	assert.Equal(t, res.Data.PayedAmout, 100000)
	assert.Equal(t, res.Data.Status, getByIDPending)
	assert.Equal(t, res.Data.Note, "ini test")
	assert.Equal(t, res.Data.Evidence, "evidence.png")
	assert.NotEmpty(t, res.Data.SubmittedAt)
	assert.NotEmpty(t, res.Data.UpdatedAt)
}
