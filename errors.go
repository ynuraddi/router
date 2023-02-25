package router

import "net/http"

var (
	notFound         func(w http.ResponseWriter) = notFoundDefault
	methodNotAllowed func(w http.ResponseWriter) = methodNotAllowedDefault
)

func (mux *Router) NotFound(f func(w http.ResponseWriter)) {
	notFound = f
}

func (mux *Router) MethodNotAllowed(f func(w http.ResponseWriter)) {
	methodNotAllowed = f
}

func notFoundDefault(w http.ResponseWriter) {
	w.Write([]byte("not found\n"))
}

func methodNotAllowedDefault(w http.ResponseWriter) {
	w.Write([]byte("method not allowed\n"))
}
