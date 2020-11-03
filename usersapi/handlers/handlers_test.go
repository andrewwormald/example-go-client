package handlers_test

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"example/users"
	"example/users/client"
	"example/usersapi/handlers"
)

// TestUserHandler uses a mock implementation of the client in order to test that the handler does what it needs to
// using the client.
func TestUserHandler(t *testing.T) {
	u := &users.User{
		ID:    123456789,
		Email: "test@test.com",
	}

	mockClient := client.NewMock(u, nil)
	handlerFunc := handlers.UserHandler(mockClient)

	req := httptest.NewRequest("GET", "http://example.com/user?email=test@test.com", nil)
	w := httptest.NewRecorder()
	handlerFunc(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
}
