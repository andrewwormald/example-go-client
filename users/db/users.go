package db

import (
	"database/sql"
	"example/users"
)

func UserWithEmail(dbc *sql.DB, email string) (*users.User, error) {
	// TODO: Query db and replace results with mock
	return &users.User{
		ID:    1,
		Email: "andrewwormald@test.com",
	}, nil
}
