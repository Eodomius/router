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

// GET
func (ro Router) Get(route string, cb func(w http.ResponseWriter, r *http.Request)){
	ro.routes[route] = cb
	fmt.Println("GET: " + route)
}

// POST
func (ro Router) Post(route string, cb func(w http.ResponseWriter, r *http.Request)){
	ro.routes[route] = cb
	fmt.Println("POST: " + route)
}

// PATCH
func (ro Router) Patch(route string, cb func(w http.ResponseWriter, r *http.Request)){
	ro.routes[route] = cb
	fmt.Println("PATCH: " + route)

}

// PUT
func (ro Router) Put(route string, cb func(w http.ResponseWriter, r *http.Request)){
	ro.routes[route] = cb
	fmt.Println("PUT: " + route)
}

// DELETE
func (ro Router) Delete(route string, cb func(w http.ResponseWriter, r *http.Request)){
	ro.routes[route] = cb
	fmt.Println("DELETE: " + route)
}
