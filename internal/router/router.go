package router

import (
	"net/http"
	"regexp"
	"strings"
)

// Handler ..
type Handler func(*Context)

// Route ..
type Route struct {
	Method  string
	Pattern *regexp.Regexp
	Handler Handler
}

// Router ..
type Router struct {
	Routes       []Route
	DefaultRoute Handler
}

// NewRouter ..
func NewRouter() *Router {
	return &Router{
		// DefaultRoute: handlers.ErrorHandler(),
	}
}

// addRoute ..
func (r *Router) addRoute(method, path string, handler Handler) {
	re := regexp.MustCompile("^" + path + "$")
	route := Route{Method: method, Pattern: re, Handler: handler}
	r.Routes = append(r.Routes, route)
}

// GET ..
func (r *Router) GET(path string, handler Handler) {
	r.addRoute("GET", path, handler)
}

// POST ..
func (r *Router) POST(path string, handler Handler) {
	r.addRoute("POST", path, handler)
}

// PUT ..
func (r *Router) PUT(path string, handler Handler) {
	r.addRoute("PUT", path, handler)
}

// DELETE ..
func (r *Router) DELETE(path string, handler Handler) {
	r.addRoute("DELETE", path, handler)
}

// Serve ..
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := &Context{ResponseWriter: w, Request: req}
	var allow []string
	for _, route := range r.Routes {
		matches := route.Pattern.FindStringSubmatch(req.URL.Path)
		if len(matches) > 0 {
			if req.Method != route.Method {
				allow = append(allow, route.Method)
				continue
			}

			if len(matches) > 1 {
				ctx.Params = matches[1:]
			}

			route.Handler(ctx)
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if req.URL.Path == "/" {
		http.Redirect(ctx.ResponseWriter, ctx.Request, "/all", 301)
	}
	http.NotFound(w, req)
	//defaultroute
}
