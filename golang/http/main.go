package main

import (
	"net/http"
)

func main() {
	h := HelloHandler2(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello2"))
	})

	// HandleはHandlerを受け取る
	// HandlerはServeHTTPをもつinterface
	// HandlerFuncはServeHTTPを持ってる
	http.Handle("/hello", &HelloHandler{})                                                 // ServeHTTPを持った構造体を自分で作った
	http.Handle("/hello2", h)                                                              // 自分で作った型にcastさせた
	http.Handle("/hello3", http.HandlerFunc(hello3))                                       // http.HandlerFuncにcastさせた
	http.Handle("/hello4", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // http.HandlerFuncに無名関数をcastさせた
		_, _ = w.Write([]byte("Hello4"))
	}))
	http.HandleFunc("/hello5", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello5"))
	})
}

type HelloHandler struct{}
type HelloHandler2 func(http.ResponseWriter, *http.Request) // http.HandlerFuncと同じ

func (HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello"))
}

func (f HelloHandler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func hello3(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello3"))
}
