package router

import (
	"net/http"
)

/**
* Resolve routes when a request is received
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
func resolveParams(route *Route, req *http.Request, result *Result) {
		params := route.PathRegex.FindStringSubmatch(req.URL.Path)
		params = params[1:]
		var paramsNumber = len(route.ParamsNames)
		for i, param := range params {
			if i > (paramsNumber-1) {
				break
			}
			result.Params[route.ParamsNames[i]] = param
		}
}