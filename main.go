package main

import (
	"net/http"

	"example/users/client"
	userapi "example/usersapi/routes"
)

func main() {
	logicalClient := client.NewLogical(nil)
	userapi.Register(logicalClient)

	// Can swap the logical for the http implementations of the client
	//httpClient := client.New("localhost", ":8080")
	//userapi.Register(httpClient)

	http.ListenAndServe(":5000", nil)
}
