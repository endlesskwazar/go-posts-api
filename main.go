package main

import (
	"context"
	"go-cource-api/infrustructure/persistence"
	"go-cource-api/interfaces"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Route struct {
	method string
	url string
	handler func(w http.ResponseWriter, r *http.Request)
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

type ctxKey struct{}


func main() {
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")

	println("host")
	println(host)

	services, err := persistence.NewRepositories(user, password, port, host, dbname)
	if err != nil {
		panic(err)
	}
	services.Automigrate()

	posts := interfaces.NewPosts(services.Post)

	var routes = []route{
		newRoute("GET", "/api/posts", posts.List),
		newRoute("POST", "/api/posts", posts.Save),
	}

	serve := func(w http.ResponseWriter, r *http.Request) {
		var allow []string
		for _, route := range routes {
			matches := route.regex.FindStringSubmatch(r.URL.Path)
			if len(matches) > 0 {
				if r.Method != route.method {
					allow = append(allow, route.method)
					continue
				}
				ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
				route.handler(w, r.WithContext(ctx))
				return
			}
		}
		if len(allow) > 0 {
			w.Header().Set("Allow", strings.Join(allow, ", "))
			http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
			return
		}
		http.NotFound(w, r)
	}

	http.HandleFunc("/", serve)

	error := http.ListenAndServe(":8000", nil)

	if error != nil {
		log.Fatal(error)
	}
}
