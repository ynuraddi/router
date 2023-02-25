package router

import (
	"net/http"
)

type Router struct {
	methodHandler  map[string]*http.ServeMux
	patternMethods map[string]map[string]bool
}

func NewRouter() *Router {
	return &Router{
		methodHandler:  make(map[string]*http.ServeMux),
		patternMethods: make(map[string]map[string]bool),
	}
}

func (mux *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serv, ok := mux.methodHandler[r.Method]
	if !ok {
		methodNotAllowed(w)
		return
	}

	handler, pattern := serv.Handler(r)
	patternMatch := mux.longestMatch(r)

	if pattern != r.URL.Path {
		notFound(w)
		return
	}

	if !mux.patternMethods[patternMatch][r.Method] {
		methodNotAllowed(w)
		return
	}

	handler.ServeHTTP(w, r)
}

func (mux *Router) longestMatch(r *http.Request) (s string) {
	for _, serv := range mux.methodHandler {
		if _, pattern := serv.Handler(r); pattern > s {
			s = pattern
		}
	}
	return s
}
