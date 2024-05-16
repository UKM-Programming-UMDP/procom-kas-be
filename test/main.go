package test

import (
	"backend/api/balance"
	"backend/api/balanceHistory"
	"backend/api/fileUpload"
	"backend/api/financialRequest"
	"backend/api/kasSubmission"
	"backend/api/month"
	"backend/api/user"
	config_database "backend/config/database"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TestConfig struct {
	Rec       *httptest.ResponseRecorder
	ServeHTTP func(w http.ResponseWriter, req *http.Request)
	Db        *gorm.DB
}

type CommonResponse struct {
	Status     bool   `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func InitTest() TestConfig {
	router := InitTestRouter()
	db := config_database.InitDB()
	InitTestRoutes(router, db)
	rec := httptest.NewRecorder()

	testConfig := TestConfig{
		Rec:       rec,
		ServeHTTP: router.ServeHTTP,
		Db:        db,
	}

	ClearAllDummy(testConfig)
	return testConfig
}

func InitTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	return router
}

func InitTestRoutes(router *gin.Engine, db *gorm.DB) {
	month.Routes(router, db)
	user.Routes(router, db)
	kasSubmission.Routes(router, db)
	fileUpload.Routes(router)
	balance.Routes(router, db)
	balanceHistory.Routes(router, db)
	financialRequest.Routes(router, db)
}

func Request(method, url string, body interface{}) *http.Request {
	jsonValue, _ := json.Marshal(body)
	bufferedJson := bytes.NewBuffer(jsonValue)
	req, _ := http.NewRequest(method, url, bufferedJson)
	return req
}

func ParseResponse(rec *httptest.ResponseRecorder, response interface{}) error {
	err := json.Unmarshal(rec.Body.Bytes(), response)
	return err
}

func WithParam(baseURL string, params interface{}) string {
	t := reflect.TypeOf(params)
	v := reflect.ValueOf(params)
	var parsedParams []string

	for i := 0; i < t.NumField(); i++ {
		formTag := t.Field(i).Tag.Get("form")
		value := v.Field(i)
		paramValue := fmt.Sprintf("%v", value.Interface())
		parsedParams = append(parsedParams, fmt.Sprintf("%s=%s", formTag, paramValue))
	}

	return baseURL + "?" + strings.Join(parsedParams, "&")
}
