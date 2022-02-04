package router

// import (
// 	"context"
// 	"net/http"
// 	"regexp"
// 	"strings"
// )

// // var routes = []route{
// // 	newRoute("GET", "/", home),
// // 	newRoute("GET", "/contact", contact),
// // 	newRoute("GET", "/api/widgets", apiGetWidgets),
// // 	newRoute("POST", "/api/widgets", apiCreateWidget),
// // 	newRoute("POST", "/api/widgets/([^/]+)", apiUpdateWidget),
// // 	newRoute("POST", "/api/widgets/([^/]+)/parts", apiCreateWidgetPart),
// // 	newRoute("POST", "/api/widgets/([^/]+)/parts/([0-9]+)/update", apiUpdateWidgetPart),
// // 	newRoute("POST", "/api/widgets/([^/]+)/parts/([0-9]+)/delete", apiDeleteWidgetPart),
// // 	newRoute("GET", "/([^/]+)", widget),
// // 	newRoute("GET", "/([^/]+)/admin", widgetAdmin),
// // 	newRoute("POST", "/([^/]+)/image", widgetImage),
// // }

// // NewRoute ..
// func (r *Rerouter) NewRoute(method, pattern string, handler http.HandlerFunc) {
// 	r.routes = append(r.routes, route{method, regexp.MustCompile("^" + pattern + "$"), handler})
// }

// type route struct {
// 	method  string
// 	regex   *regexp.Regexp
// 	handler http.HandlerFunc
// }

// // Rerouter ..
// type Rerouter struct {
// 	routes []route
// }

// // NewReRouter ..
// func NewReRouter() *Rerouter {
// 	return &Rerouter{}
// }

// // Serve ..
// func (r *Rerouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	var allow []string
// 	for _, route := range r.routes {
// 		matches := route.regex.FindStringSubmatch(req.URL.Path)
// 		if len(matches) > 0 {
// 			if req.Method != route.method {
// 				allow = append(allow, route.method)
// 				continue
// 			}
// 			ctx := context.WithValue(req.Context(), CtxKey{}, matches[1:])
// 			route.handler(w, req.WithContext(ctx))
// 			return
// 		}
// 	}
// 	if len(allow) > 0 {
// 		w.Header().Set("Allow", strings.Join(allow, ", "))
// 		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
// 	http.NotFound(w, req)
// }

// // CtxKey ..
// type CtxKey struct{}

// // GetField ..
// func GetField(r *http.Request, index int) string {
// 	fields := r.Context().Value(CtxKey{}).([]string)
// 	return fields[index]
// }
