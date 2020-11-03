package client

import (
	"database/sql"

	"example/users"
	"example/users/db"
)

func NewLogical(dbc *sql.DB) Client {
	return &logicalImpl{
		dbc: dbc,
	}
}

type logicalImpl struct {
	dbc *sql.DB
}

func (c *logicalImpl) UserWithEmail(email string) (*users.User, error) {
	return db.UserWithEmail(c.dbc, email)
}
