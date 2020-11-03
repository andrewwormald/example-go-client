package routes

import (
	"net/http"

	"example/users/client"
	"example/usersapi/handlers"
)

func Register(usersClient client.Client) {
	http.HandleFunc("/user", handlers.UserHandler(usersClient))
}
