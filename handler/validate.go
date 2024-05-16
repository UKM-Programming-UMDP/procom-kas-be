package handler

import (
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func BindAndValidate(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		if err == io.EOF {
			Error(c, http.StatusBadRequest, "Invalid JSON data: unexpected end of JSON input")
			return true
		}
		Error(c, http.StatusBadRequest, err.Error())
		return true
	}
	if err := requestValidator(c, req, "body"); err != nil {
		return true
	}
	return false
}

func BindParamAndValidate(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBindQuery(req); err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return true
	}
	if err := requestValidator(c, req, "param"); err != nil {
		return true
	}
	return false
}

func QueryValidator(query *gorm.DB, c *gin.Context, count bool) (bool, int64) {
	if query.Error != nil {
		Error(c, http.StatusInternalServerError, query.Error.Error())
		return false, -1
	}

	if !count {
		return true, -1
	}

	var result int64
	if query.Count(&result); result == 0 {
		return true, 0
	}

	return true, result
}

var tagMessages = map[string]string{
	"required":       "This field is required",
	"required-param": "This parameter is required",
	"email":          "Invalid email",
	"max":            "Exceeds maximum length",
	"min":            "Exceeds minimum length",
	"len":            "Invalid length",
}

func parseTagMessage(tag string) string {
	if message, ok := tagMessages[tag]; ok {
		return message
	}
	return tag
}

func requestValidator(c *gin.Context, req interface{}, validatorType string) error {
	var modelTag string
	var errorMessage string

	if validatorType == "body" {
		modelTag = "json"
		errorMessage = "Invalid request body"
	} else if validatorType == "param" {
		modelTag = "form"
		errorMessage = "Invalid query parameter"
	} else {
		panic("Invalid validator type")
	}

	val := validator.New()
	val.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get(modelTag), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	if err := val.Struct(req); err != nil {
		var valErrors validator.ValidationErrors
		if errors.As(err, &valErrors) {
			errors := make([]ApiError, len(valErrors))
			for i, valError := range valErrors {
				errorTag := valError.Tag()
				if errorTag == "required" && validatorType == "param" {
					errorTag = "required-param"
				}
				errors[i] = ApiError{valError.Field(), parseTagMessage(errorTag)}
			}
			Error(c, http.StatusBadRequest, errorMessage, errors...)
		}
		return err
	}
	return nil
}
