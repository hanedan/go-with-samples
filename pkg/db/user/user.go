package db

import "database/sql"

type UserDB struct {
	db *sql.DB
}

func NewUserDB(db *sql.DB) *UserDB {
	return &UserDB{db: db}
}

type User struct {
	UserID   int    `json:"id"`
	Name     string `json:"name" validate:"required,min=2,max=255,nodots"`
	LastName string `json:"last_name" validate:"required,min=2,max=255"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Mobile   string `json:"mobile" validate:"required,e164"`
	Birthday string `json:"birthday" validate:"required,datetime=2006-01-02"`
}
