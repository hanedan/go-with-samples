package model

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var v validator.Validate

func init() {
	v = *validator.New()

	// Register custom validators sample
	err := v.RegisterValidation("nodots", func(fl validator.FieldLevel) bool {
		t, ok := fl.Field().Interface().(string)
		if !ok {
			return false
		}
		if err := ValidateNoDots(t); err != nil {
			return false
		}
		return true
	})
	if err != nil {
		panic(err)
	}
}

func ValidateNoDots(s string) error {
	if strings.Contains(s, ".") {
		return fmt.Errorf("string %s contains dots", s)
	}
	return nil
}

func InvalidRequest(c *gin.Context, expMethod string) bool {
	if c.Request.Method != expMethod {
		c.JSON(http.StatusBadRequest, NewResponse(WithError(http.StatusBadRequest, ErrInvalidRequestMethod)))
		return true
	}

	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, NewResponse(WithError(http.StatusBadRequest, ErrEmptyRequestBody)))
		return true
	}
	return false
}
