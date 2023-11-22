package setup

import (
	"context"
	"database/sql"
)

func CreateUserTable(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, `
	CREATE TABLE IF NOT EXISTS Users (
		UserID SERIAL PRIMARY KEY,
		Name VARCHAR(255) NOT NULL,
		LastName VARCHAR(255) NOT NULL,
		Email VARCHAR(255) NOT NULL,
		Mobile VARCHAR(16) NOT NULL,
		Birthday DATE NOT NULL,
		UNIQUE (Email)
	);`)
	if err != nil {
		return err
	}
	return nil
}
