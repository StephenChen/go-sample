package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"web/middleware"
	"web/router"
	"web/validator"
)

func main() {
	// router
	router.Router()

	// middleware
	middleware.Middleware()

	// validator
	validator.Validate()
	validator.ValidateSelf()

	http.Handle("/ccc", new(appHandler))
	http.Handle("/eee", http.HandlerFunc(echo))
	http.Handle("/ddd", something(http.HandlerFunc(echo)))
	http.HandleFunc("/", echo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("echo error"))
		return
	}

	writeLen, err := w.Write(msg)
	if err != nil || writeLen != len(msg) {
		log.Println(err, "write len:", writeLen)
	}
}

type appHandler struct {
}

func (a *appHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("its in ccc"))
}

func something(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do something
		next.ServeHTTP(w, r)
		// do something
	})
}

func something1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do something
		next.ServeHTTP(w, r)
		// do something
	})
}

func something2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do something
		next.ServeHTTP(w, r)
		// do something
	})
}
