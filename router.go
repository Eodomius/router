package router

import (
	"fmt"
	"net/http"
	"regexp"
)



func New() *Router {
	return &Router{
		routes: make(map[string]Route),
		middlewares: make([]func(w http.ResponseWriter, r *http.Request), 0),
		paramsRegex: regexp.MustCompile(`({[^/]+})`),
	}
}


func (ro Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	req.URL.Path = normalizePath(req.URL.Path)
		route, exist := resolveRoute(ro, req)
		if !exist {
			return
		}
		
		var result = Result{
			Params: make(map[string]string),
		}
		resolveParams(&route, req, &result)
		for _, middleware := range ro.middlewares {
			middleware(res, req)
		}
		route.Execute(res, req, &result)
}


/* It's a method that adds a route to the router. */
func (ro Router) HandleRoute(path, method string, cb func(w http.ResponseWriter, r *http.Request, route *Result)){
	route := Route{
		Path: path,
		Method: method,
		Execute: cb,
		ParamsNames: []string{},
	}
	paramsNames := ro.paramsRegex.FindAllString(path, -1)
	for i, paramName := range paramsNames {
		paramsNames[i] = regexp.MustCompile("{|}").ReplaceAllString(paramName, "")
	}
	route.ParamsNames = paramsNames
	var replacedRoute = ro.paramsRegex.ReplaceAllString(path, `([^/]+)`)
	replacedRoute = "^" + replacedRoute + "$";
	route.PathRegex = regexp.MustCompile(replacedRoute)
	ro.routes[method + ": " + path] = route
	fmt.Println(method + ": " + route.Path)
}



/****************** Methods handlers ******************/

// Add middleware
func (ro *Router) Use(middleware func(w http.ResponseWriter, r *http.Request)){
	ro.middlewares = append(ro.middlewares, middleware)
}

// GET
func (ro Router) Get(path string, cb func(w http.ResponseWriter, r *http.Request, route *Result)){
	ro.HandleRoute(path, "GET", cb)
}

// HEAD
func (ro Router) Head(path string, cb func(w http.ResponseWriter, r *http.Request, route *Result)){
	ro.HandleRoute(path, "HEAD", cb)
}

// POST
func (ro Router) Post(path string, cb func(w http.ResponseWriter, r *http.Request, route *Result)){
	ro.HandleRoute(path, "POST", cb)
}

// PATCH
func (ro Router) Patch(path string, cb func(w http.ResponseWriter, r *http.Request, route *Result)){
	ro.HandleRoute(path, "PATCH", cb)
}

// PUT
func (ro Router) Put(path string, cb func(w http.ResponseWriter, r *http.Request, route *Result)){
	ro.HandleRoute(path, "PUT", cb)
}

// DELETE
func (ro Router) Delete(path string, cb func(w http.ResponseWriter, r *http.Request, route *Result)){
	ro.HandleRoute(path, "DELETE", cb)
}
