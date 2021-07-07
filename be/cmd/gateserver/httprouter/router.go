package httprouter

import (
	"context"
	"net/http"
)

// The handler for the router item
type RouteHandler func(http.ResponseWriter, *http.Request, Params)

// The parameter pair of the request
type Param struct {
	Key   string
	Value string
}

// The parameter pair collection of the request
type Params []Param

func (ps Params) Get(name string) string {
	for i := range ps {
		if ps[i].Key == name {
			return ps[i].Value
		}
	}
	return ""
}

// The parameter keys
type paramsKey struct{}

// The parameter keys
var ParamsKey = paramsKey{}

// Parses the parameter from the context
func ParamsFromContext(ctx context.Context) Params {
	p, _ := ctx.Value(ParamsKey).(Params)
	return p
}

// The router object
type Router struct {
	trees        map[string]*node
	PanicHandler func(http.ResponseWriter, *http.Request, interface{})
}

// Checks whether the Router conforms the http.Handler
var _ http.Handler = New()

// Constructor
func New() *Router {
	return &Router{}
}

// Adds a route handler for path with GET method
func (r *Router) GET(path string, handle RouteHandler) {
	r.AddRouteHandler(http.MethodGet, path, handle)
}

// Adds a route handler for path with POST method
func (r *Router) POST(path string, handle RouteHandler) {
	r.AddRouteHandler(http.MethodPost, path, handle)
}

// Adds a route handler for path with PUT method
func (r *Router) PUT(path string, handle RouteHandler) {
	r.AddRouteHandler(http.MethodPut, path, handle)
}

// Adds a route handler for path with PUT method
func (r *Router) OPTIONS(path string, handle RouteHandler) {
	r.AddRouteHandler(http.MethodOptions, path, handle)
}

// Adds a route handler for path with method
func (r *Router) AddRouteHandler(method, path string, handle RouteHandler) {
	// if path is empty or is only one char but not "/", panic
	if len(path) < 1 || path[0] != '/' {
		panic("path must begin with '/' in path '" + path + "'")
	}

	if r.trees == nil {
		// create the trie
		r.trees = make(map[string]*node)
	}

	// get the node for current method
	root := r.trees[method]
	if root == nil {
		root = new(node)
		r.trees[method] = root
	}

	// add the route to the node
	root.addRoute(path, handle)
}

// recv
func (r *Router) recv(w http.ResponseWriter, req *http.Request) {
	if rcv := recover(); rcv != nil {
		r.PanicHandler(w, req, rcv)
	}
}

// Enables the cors for all request to this server
func enableCors(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// Implements the global http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if r.PanicHandler != nil {
		defer r.recv(w, req)
	}

	// enable cors
	enableCors(w, req)

	// return for optins
	if req.Method == http.MethodOptions {
		return
	}

	// get the request path
	path := req.URL.Path

	// find the handler for the request
	if root := r.trees[req.Method]; root != nil {
		if handler, ps, _ := root.getValue(path); handler != nil {
			handler(w, req, ps)
			return
		}
	}

	// 404 Not found
	http.NotFound(w, req)
}
