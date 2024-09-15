package auth

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"webengine/core"
)

type User struct {
	Name string `json:"name"`
}

func Login(a core.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a.Database().KvSet(`{"name": "BENNY", "loggedIn": true, "key": "100"}`)
		w.Write([]byte("Login page"))
	}
}

func UserPage(a core.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := a.Query().GetUserByEmail(r.Context(), "BENNY")

		user, err := a.Database().KvGet("100")
		if err != nil {
			w.Write([]byte("User not found"))
			return
		}

		var u User
		err = json.Unmarshal([]byte(user), &u)

		if err != nil {
			return
		}

		w.Write([]byte("Hello " + u.Name))
	}
}

func RegisterRoutes(a core.Application) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/login", Login(a))
		r.Get("/user", UserPage(a))
	}
}
