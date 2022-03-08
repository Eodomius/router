package router

import (
	"fmt"
	"net/http"
	"regexp"
)
type Route struct {
	Method string
	Path string
	PathRegex *regexp.Regexp
	Params map[string]string
	ParamsNames []string
	Execute func(w http.ResponseWriter, r *http.Request, route Route)
}
type Router struct {
	routes map[string]Route
	paramsRegex *regexp.Regexp
}

func New() Router {
	return Router{
		routes: make(map[string]Route),
		paramsRegex: regexp.MustCompile(`({[^/]+})`),
	}
}

/**
* Resolve routes
* @param {Router} Router - The router
* @param {http.Request} req - The request
* @return {Route} - The route
* @return {bool} - If the route exist
*/
func resolveRoute(ro Router, r *http.Request) (Route, bool) {
	for _, route := range ro.routes {
			if route.Method == r.Method {
				if route.PathRegex.MatchString(r.URL.Path) {
					return route , true
				}
			}
		}
		return Route{}, false
}

/**
* Resolve routes and update Route.Params
* @param {*Route} Route - The route called
* @param {http.Request} req - The request
* @return {[]string} - The params
*/
func resolveParams(route *Route, req *http.Request) {
		params := route.PathRegex.FindStringSubmatch(req.URL.Path)
		params = params[1:]
		for i, param := range params {
			route.Params[route.ParamsNames[i]] = param
		}
}
func (ro Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {

		route, exist := resolveRoute(ro, req)
		if !exist {
			return
		}
		resolveParams(&route, req)
		route.Execute(res, req, route)
}

/* Methods functions */

func (ro Router) handle(path, method string, cb func(w http.ResponseWriter, r *http.Request, route Route)){
	route := Route{
		Path: path,
		Method: method,
		Execute: cb,
		Params: make(map[string]string),
		ParamsNames: []string{},
	}
	// Get params names
	paramsNames := ro.paramsRegex.FindAllString(path, -1)
	route.ParamsNames = paramsNames
	// Replace {paramName} by ([^/]+)
	var replacedRoute = ro.paramsRegex.ReplaceAllString(path, `([^/]+)`)
	route.PathRegex = regexp.MustCompile(replacedRoute)
	ro.routes[method + ": " + path] = route
	fmt.Println(method + ": " + route.Path)
}

// GET
func (ro Router) Get(path string, cb func(w http.ResponseWriter, r *http.Request, route Route)){
	ro.handle(path, "GET", cb)
}

// POST
func (ro Router) Post(path string, cb func(w http.ResponseWriter, r *http.Request, route Route)){
	ro.handle(path, "POST", cb)
}

// PATCH
func (ro Router) Patch(path string, cb func(w http.ResponseWriter, r *http.Request, route Route)){
	ro.handle(path, "PATCH", cb)
}

// PUT
func (ro Router) Put(path string, cb func(w http.ResponseWriter, r *http.Request, route Route)){
	ro.handle(path, "PUT", cb)
}

// DELETE
func (ro Router) Delete(path string, cb func(w http.ResponseWriter, r *http.Request, route Route)){
	ro.handle(path, "DELETE", cb)
}
