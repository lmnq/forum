package router

import (
	"fmt"
	"net/http"
	"regexp"
)

// Route ..
type Route struct {
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

// Router ..
type Router struct {
	routes []Route
}

// NewRouter ..
func NewRouter() *Router {
	return &Router{}
}

// AddRoute ..
func (r *Router) AddRoute(method, path string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{Method: method, Pattern: path, Handler: handler})
}

// GET ..
func (r *Router) GET(path string, handler http.HandlerFunc) {
	r.AddRoute("GET", path, handler)
}

// POST ..
func (r *Router) POST(path string, handler http.HandlerFunc) {
	r.AddRoute("POST", path, handler)
}

// PUT ..
func (r *Router) PUT(path string, handler http.HandlerFunc) {
	r.AddRoute("PUT", path, handler)
}

// DELETE ..
func (r *Router) DELETE(path string, handler http.HandlerFunc) {
	r.AddRoute("DELETE", path, handler)
}

func (r *Router) getHandler(method, path string) http.HandlerFunc {
	for _, route := range r.routes {
		re := regexp.MustCompile(route.Pattern)
		if route.Method == method && re.MatchString(path) {
			// match := re.FindStringSubmatch(path)
			// if match != nil {
			return route.Handler
			// }
		}
	}
	return http.NotFoundHandler().ServeHTTP
}

// ServeHTTP ..
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method
	fmt.Println(method, path)

	handler := r.getHandler(method, path)

	// handler middleware

	handler(w, req)
}
