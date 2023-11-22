package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	u "go-with-samples/pkg/db/user"
)

func TestCreateHandlerError(t *testing.T) {
	db, _, err := sqlmock.New()
	require.NoError(t, err, "failed to create mock database connection: %v", err)
	t.Cleanup(func() {
		db.Close()
	})

	router := gin.Default()
	userAPI := NewUserAPI(u.NewUserDB(db))
	uri := "/v1/library-api/users/create"
	router.POST(uri, userAPI.CreateHandler)

	invalidRequests := map[string]struct {
		request     any
		expectedMsg string
	}{
		"ReqID must be a UUID. A Request with Non-UUID ReqID must return BadRequest": {
			request: CreateUserReq{
				ReqID: "120",
				User: u.User{
					Name:     "Name",
					LastName: "LastName",
					Email:    "valid@email.com",
					Mobile:   "+905001001010",
					Birthday: "2000-01-01",
				},
			},
			expectedMsg: ErrValidationFailed,
		},
		"ReqID is valid, but User data is not. Dateformat doesn't match, must return BadRequest": {
			request: CreateUserReq{
				ReqID: "120",
				User: u.User{
					Name:     "Name",
					LastName: "LastName",
					Email:    "valid@email.com",
					Mobile:   "+905001001010",
					Birthday: "01-01-2001",
				},
			},
			expectedMsg: ErrValidationFailed,
		},
		"The request type is not kind of CreateUserRequest, therefore must return BadRequest": {
			request:     "an invalid request",
			expectedMsg: ErrRequestBodyParseFailed,
		},
	}

	for desc, tt := range invalidRequests {
		t.Run(desc, func(t *testing.T) {
			jsonBody, err := json.Marshal(tt.request)
			assert.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(jsonBody))
			assert.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
			var resp Response
			err = json.Unmarshal(w.Body.Bytes(), &resp)
			assert.NoError(t, err)
			assert.True(t, strings.Contains(resp.Error, tt.expectedMsg))
		})
	}

	t.Run("request with empty body, must return BadRequest", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, uri, nil)
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var resp Response
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.True(t, strings.Contains(resp.Error, ErrEmptyRequestBody))
	})
}

func TestCreateHandlerDuplicateKeyError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err, "failed to create mock database connection: %v", err)
	defer db.Close()

	router := gin.Default()
	userAPI := NewUserAPI(u.NewUserDB(db))
	uri := "/v1/library-api/users/create"
	router.POST(uri, userAPI.CreateHandler)

	request := CreateUserReq{
		ReqID: "649fb032-f692-419f-b522-d11c87e40468",
		User: u.User{
			Name:     "Name",
			LastName: "LastName",
			Email:    "valid@email.com",
			Mobile:   "+905001001010",
			Birthday: "2000-01-01",
		},
	}

	mock.ExpectPrepare("INSERT INTO users (.+) VALUES (.+)").
		ExpectExec().WithArgs(request.User.Name, request.User.LastName, request.User.Email, request.User.Mobile, request.User.Birthday).
		WillReturnError(fmt.Errorf(ErrCouldntCreateUser + `pq: duplicate key value violates unique constraint "users_email_key"`))

	jsonBody, err := json.Marshal(request)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(jsonBody))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var resp Response
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.True(t, strings.Contains(resp.Error, ErrCouldntCreateUser), "must return duplicate key error")

	err = mock.ExpectationsWereMet()
	require.NoError(t, err, "there were unfulfilled expectations: %v", err)
}
