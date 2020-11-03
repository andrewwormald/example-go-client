package client

import (
	"encoding/json"
	"net/http"

	"example/users"
)

func New(domain, port string) Client {
	return &httpImpl{
		domain: domain,
		port: port,
	}
}

type httpImpl struct {
	domain string
	port string
}

func (c *httpImpl) UserWithEmail(email string) (*users.User, error) {
	url := c.domain + c.port
	resp, err := http.Get(url + "/user" + "?email=" + email)

	u := new(users.User)
	err = json.NewDecoder(resp.Body).Decode(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
