package http

import (
	"database/sql"
	"encoding/json"
	"example/users/db"
	"net/http"
)

func Run(port string, dbc *sql.DB) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		if email == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("email param required"))
		}

		user, err := db.UserWithEmail(dbc, email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(port, nil)
}
