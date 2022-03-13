package router

import (
	"net/http"
	"regexp"
)


type Route struct {
	Method string
	Path string
	PathRegex *regexp.Regexp
	ParamsNames []string
	Execute func(w http.ResponseWriter, r *http.Request, route *Result)
}
type Router struct {
	routes map[string]Route
	middlewares []func(w http.ResponseWriter, r *http.Request)
	paramsRegex *regexp.Regexp
}
type Result struct {
		Params map[string]string
}