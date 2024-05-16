package financialRequest_test

import (
	. "backend/test"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type createStatus string
type createPaymentType string

const (
	Approved createStatus = "Approved"
	Rejected createStatus = "Rejected"
	Pending  createStatus = "Pending"
)
const (
	Cash     createPaymentType = "Cash"
	Transfer createPaymentType = "Transfer"
)

type createPayment struct {
	Type           createPaymentType `json:"type"`
	TargetProvider string            `json:"target_provider"`
	TargetName     string            `json:"target_name"`
	TargetNumber   string            `json:"target_number"`
	Evidence       string            `json:"evidence"`
}

type createRequestBody struct {
	Amount int    `json:"amount"`
	Note   string `json:"note"`
	User   struct {
		NPM string `json:"npm"`
	} `json:"user"`
	createPayment `json:"payment"`
}

type createResponse struct {
	CommonResponse
	Data struct {
		RequestID string `json:"request_id"`
		Amount    int    `json:"amount"`
		Note      string `json:"note"`
		User      struct {
			NPM   string `json:"npm"`
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"user,omitempty"`
		Status             createStatus  `json:"status"`
		Payment            createPayment `json:"payment"`
		TransferedEvidence string        `json:"transfered_evidence"`
		CreatedAt          time.Time     `json:"created_at"`
		UpdatedAt          time.Time     `json:"updated_at"`
	} `json:"data"`
}

func TestCreateFinancialRequest(t *testing.T) {
	config := InitTest()
	CreateUserDummy(config)

	newFinreq := createRequestBody{
		Amount: 100000,
		Note:   "ini test",
		User: struct {
			NPM string `json:"npm"`
		}{
			NPM: "1928476912",
		},
		createPayment: createPayment{
			Type:           Transfer,
			TargetProvider: "Bank - TEST",
			TargetName:     "user_test",
			TargetNumber:   "1234567890",
			Evidence:       "evidence.png",
		},
	}
	req := Request("POST", "/api/financial-request", newFinreq)
	config.ServeHTTP(config.Rec, req)
	assert.Equal(t, http.StatusCreated, config.Rec.Code)
	DeleteFinancialRequestDummy(config)

	var res createResponse
	if err := ParseResponse(config.Rec, &res); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.Message, "Success creating a financial request")
	assert.Len(t, res.Data.RequestID, 5)
	assert.Equal(t, res.Data.User.NPM, "1928476912")
	assert.Equal(t, res.Data.User.Name, "user_test")
	assert.Equal(t, res.Data.User.Email, "user_test@mail.com")
	assert.Equal(t, res.Data.Amount, 100000)
	assert.Equal(t, res.Data.Note, "ini test")
	assert.Equal(t, res.Data.Status, Pending)
	assert.Equal(t, res.Data.Payment.Type, Transfer)
	assert.Equal(t, res.Data.Payment.TargetProvider, "Bank - TEST")
	assert.Equal(t, res.Data.Payment.TargetName, "user_test")
	assert.Equal(t, res.Data.Payment.TargetNumber, "1234567890")
	assert.Equal(t, res.Data.Payment.Evidence, "evidence.png")
	assert.Empty(t, res.Data.TransferedEvidence)
	assert.NotEmpty(t, res.Data.CreatedAt)
	assert.NotEmpty(t, res.Data.UpdatedAt)
}
