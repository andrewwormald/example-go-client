package client

import "example/users"

// Client is the declarative interface of the client implementation. Either use the logical implementation or the HTTP
// implementation that communicates over the wire
type Client interface {
	UserWithEmail(email string) (*users.User, error)
}


