package main

import (
	"net/http"
)

func main() {
	m := http.NewServeMux()
	h := new(HelloHandler)
	m.Handle("/index", h)

	s := http.Server{
		Addr:    ":8080",
		Handler: m,
	}
	s.ListenAndServe()
}

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello world"))
}
