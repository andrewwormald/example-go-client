package handlers

import (
	"encoding/json"
	"net/http"

	"example/users/client"
)

func UserHandler(usersClient client.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		if email == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("email param required"))
		}

		user, err := usersClient.UserWithEmail(email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
