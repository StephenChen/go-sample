package middleware

import (
	"log"
	"net/http"
)

type middleware func(handler http.Handler) http.Handler

type Router struct {
	middlewareChain []middleware
	mux             map[string]http.Handler
}

func NewRouter() *Router {
	return &Router{
		mux: make(map[string]http.Handler),
	}
}

func (r *Router) Use(m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}

func (r *Router) Add(route string, h http.Handler) {
	var mergeHandler = h

	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergeHandler = r.middlewareChain[i](mergeHandler)
	}

	r.mux[route] = mergeHandler
}

func Middleware() {
	r := NewRouter()
	r.Use(timeout)
	r.Use(logger)
	r.Use(other)
	r.Add("/", http.HandlerFunc(echo))
}

func echo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("its in middleware"))
}

func timeout(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do something
		log.Println("before timeout")
		next.ServeHTTP(w, r)
		log.Println("after timeout")
		// do something
	})
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do something
		log.Println("before logger")
		next.ServeHTTP(w, r)
		log.Println("after logger")
		// do something
	})
}

func other(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do something
		log.Println("before other")
		next.ServeHTTP(w, r)
		log.Println("after other")
		// do something
	})
}
