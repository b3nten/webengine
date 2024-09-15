package ui

import (
	"net/http"
	"webengine/core"
)

func HomeRoute(a core.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := Home("BENNY").Render(r.Context(), w)
		err = a.Database().KvSet(`{"name": "BENNY", "key": "100"}`)
		if err != nil {
			panic(err)
			return
		}
	}
}
