package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var v validator.Validate

func init() {
	v = *validator.New()
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
