package kasSubmission_test

import (
	. "backend/test"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type createRequestBody struct {
	User struct {
		NPM string `json:"npm"`
	} `json:"user"`
	PayedAmount *int    `json:"payed_amount"`
	Note        *string `json:"note"`
	Evidence    string  `json:"evidence"`
}

type createStatus string

const (
	createPending  createStatus = "Pending"
	createApproved createStatus = "Approved"
	createRejected createStatus = "Rejected"
)

type createResponse struct {
	CommonResponse
	Data struct {
		SubmissionID string `json:"submission_id"`
		User         struct {
			NPM      string `json:"npm"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			KasPayed int    `json:"kas"`
		} `json:"user"`
		PayedAmout  int          `json:"payed_amount"`
		Status      createStatus `json:"status"`
		Note        string       `json:"note"`
		Evidence    string       `json:"evidence"`
		SubmittedAt string       `json:"submitted_at"`
		UpdatedAt   string       `json:"updated_at"`
	} `json:"data"`
}

func TestCreateKasSubmission(t *testing.T) {
	config := InitTest()
	CreateUserDummy(config)

	payedAmount := 100000
	note := "ini test"
	newKasSubmission := createRequestBody{
		User: struct {
			NPM string `json:"npm"`
		}{
			NPM: "1928476912",
		},
		PayedAmount: &payedAmount,
		Note:        &note,
		Evidence:    "evidence.png",
	}
	req := Request("POST", "/api/kas", newKasSubmission)
	config.ServeHTTP(config.Rec, req)
	assert.Equal(t, http.StatusCreated, config.Rec.Code)
	DeleteKasSubmissionDummy(config)

	var res createResponse
	if err := ParseResponse(config.Rec, &res); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.Message, "Success creating a kas submission")
	assert.Len(t, res.Data.SubmissionID, 5)
	assert.Equal(t, res.Data.User.NPM, "1928476912")
	assert.Equal(t, res.Data.User.Name, "user_test")
	assert.Equal(t, res.Data.User.Email, "user_test@mail.com")
	assert.Equal(t, res.Data.User.KasPayed, 0)
	assert.Equal(t, res.Data.PayedAmout, 100000)
	assert.Equal(t, res.Data.Status, createPending)
	assert.Equal(t, res.Data.Note, "ini test")
	assert.Equal(t, res.Data.Evidence, "evidence.png")
	assert.NotEmpty(t, res.Data.SubmittedAt)
	assert.NotEmpty(t, res.Data.UpdatedAt)
}
