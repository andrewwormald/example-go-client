package client

import "example/users"

func NewMock(expectedUser *users.User, expectedErr error) Client {
	return &mock{
		expectedUser: expectedUser,
		expectedErr:  expectedErr,
	}
}

type mock struct {
	expectedUser *users.User
	expectedErr error
}

func (m *mock) UserWithEmail(email string) (*users.User, error) {
	return m.expectedUser, m.expectedErr
}
