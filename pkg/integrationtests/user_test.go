//go:build integration

// To run integration tests, use:
// go test ./... -tags=integration
// Alternatively, a file can be marked as a unit test with: //go:build unit
// However, -tags=unit must be passed to the go test command.
package integrationtests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"go-with-samples/pkg/db"
	u "go-with-samples/pkg/db/user"
)

func TestCreateUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	t.Cleanup(func() {
		cancel()
	})

	db, err := db.Connect(ctx)
	require.NoError(t, err, "failed to initialize database connection: %v", err)
	t.Cleanup(func() {
		db.Close()
	})

	userDB := u.NewUserDB(db)

	user := u.User{
		Name:     "Name",
		LastName: "LastName",
		Email:    "valid@email.com",
		Mobile:   "+905001001010",
		Birthday: "2000-01-01",
	}

	t.Cleanup(func() {
		userDB.Delete(ctx, user)
	})

	err = userDB.Create(ctx, user)
	require.NoError(t, err, "failed to create user: %v", err)

	// pq: duplicate key value violates unique constraint "users_email_key"
	err = userDB.Create(ctx, user)
	require.Error(t, err, "expected to get an error when trying to create a user with the same email.")
	require.Equal(t, `pq: duplicate key value violates unique constraint "users_email_key"`, err.Error())
}
