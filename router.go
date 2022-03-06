package router

import (
	"fmt"
	"net/http"
)

type Router struct {
	routes map[string]func(w http.ResponseWriter, r *http.Request)
}

func New()Router{
	return Router{
		routes: make(map[string]func(w http.ResponseWriter, r *http.Request)),
	}
}

func (ro Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var route = ro.routes[r.URL.Path]
		if route != nil {
			route(w, r)
		} 
	}
}

/* Methods functions */
func (ro Router) Get(route string, cb func(w http.ResponseWriter, r *http.Request)){
	ro.routes[route] = cb
	fmt.Println("GET: " + route)
}
