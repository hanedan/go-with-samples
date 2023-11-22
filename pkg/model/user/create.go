package model

import (
	"net/http"

	"github.com/gin-gonic/gin"

	db "go-with-samples/pkg/db/user"
)

type CreateUserReq struct {
	ReqID string  `json:"id" validate:"required,uuid"`
	User  db.User `json:"user" validate:"required"`
}

func (u *UserAPI) CreateHandler(c *gin.Context) {
	if InvalidRequest(c, http.MethodPost) {
		return
	}

	var req CreateUserReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, NewResponse(WithError(http.StatusBadRequest, ErrRequestBodyParseFailed+err.Error())))
		return
	} else if err = v.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, NewResponse(WithError(http.StatusBadRequest, ErrValidationFailed+err.Error())))
		return
	}

	if err := u.db.Create(c.Request.Context(), req.User); err != nil {
		c.JSON(http.StatusBadRequest, NewResponse(WithError(http.StatusBadRequest, ErrCouldntCreateUser+err.Error())))
		return
	}

	c.JSON(http.StatusOK, NewResponse(WithSuccess(http.StatusOK)))
}
