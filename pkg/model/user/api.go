package model

import (
	db "go-with-samples/pkg/db/user"
)

type UserAPI struct {
	db *db.UserDB
}

func NewUserAPI(db *db.UserDB) *UserAPI {
	return &UserAPI{db: db}
}
