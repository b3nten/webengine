package main

import (
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"net/http/httputil"
	"net/url"
	"webengine/auth"
	"webengine/core"
	"webengine/ui"
)

var isProd bool

var port int

func init() {
	flag.BoolVar(&isProd, "prod", false, "run in production mode")
	flag.IntVar(&port, "port", 8001, "port to run the server on")
	flag.Parse()
}

func getEnvironment() core.Version {
	if isProd {
		return core.VersionProd
	}
	return core.VersionDev
}

func serveWeb(r *chi.Mux) {
	if env := getEnvironment(); env == core.VersionProd {
		r.Handle("/*", http.FileServer(http.Dir("./web/dist")))
	} else {
		viteURL, err := url.Parse("http://localhost:5173")
		if err != nil {
			panic(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(viteURL)
		r.NotFound(func(w http.ResponseWriter, r *http.Request) {
			proxy.ServeHTTP(w, r)
		})
	}
}

func main() {

	r := chi.NewRouter()

	app, err := core.NewApplication(getEnvironment())

	if err != nil {
		panic(err)
	}

	r.Use(middleware.Logger)

	r.Get("/hello", ui.HomeRoute(app))

	r.Route("/auth", auth.RegisterRoutes(app))

	serveWeb(r)

	fmt.Printf("Server is running on port %d\n", port)

	var prefix string
	if isProd {
		prefix = ""
	} else {
		prefix = "127.0.0.1"
	}

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", prefix, port), r); err != nil {
		panic(err)
	}
}
