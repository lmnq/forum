package router

import "net/http"

// Context ..
type Context struct {
	http.ResponseWriter
	*http.Request
	Params []string
}